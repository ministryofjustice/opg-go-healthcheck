package gohealthcheck

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestDefaultHc(t *testing.T) {
	hc := defaultHc()
	if hc == nil {
		t.Fatal("default hc is nil")
	}
	if hc.logger == nil {
		t.Error("default logger is nil")
	}
	if hc.exit == nil {
		t.Error("exit is nil")
	}
	if hc.run {
		t.Error("should not run without -hc flag")
	}
}

func TestHealthCheckRegister(t *testing.T) {
	tests := []struct {
		run          bool
		hcStatusCode int
		wantExitCode int
		wantInLog    string
	}{
		{
			true,
			200,
			0,
			"OK",
		},
		{
			true,
			500,
			0,
			"FAIL",
		},
		{
			false,
			500,
			666,
			"",
		},
	}

	initArgs := os.Args

	for _, test := range tests {
		hc := new(healthCheck)
		buf := new(bytes.Buffer)
		hc.logger = log.New(buf, "", log.LstdFlags)
		retCode := 666
		hc.exit = func(code int) {
			retCode = code
		}

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(test.hcStatusCode)
		}))

		hc.run = test.run

		hc.register(s.URL)

		if retCode != test.wantExitCode {
			t.Errorf("Unexpected exit code. Want %d, got %d", test.wantExitCode, retCode)
		}

		if !strings.Contains(buf.String(), test.wantInLog) {
			t.Errorf("Log does not contain: %s", test.wantInLog)
		}

		s.Close()
		os.Args = initArgs
	}
}
