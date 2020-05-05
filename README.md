# opg-go-healthcheck

Module that adds a health-check utility to your Go app.

Inspired by https://medium.com/google-cloud/dockerfile-go-healthchecks-k8s-9a87d5c5b4cb

### Usage

Install as dependency

```
go get github.com/ministryofjustice/opg-go-healthcheck
```

Register a health-check address at the top of your main() function

```go
package main

import (
    "github.com/ministryofjustice/opg-go-healthcheck"
)

func main() {
    gohealthcheck.Register("http://localhost:8000/health-check")

    // ...
}
```

Run the health-check by adding a `-hc` flag to your executable

```
./your-go-build -hc
``` 
