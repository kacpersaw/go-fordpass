# go-fordpass
A Go client library for the FordPass API.

Project based on [fordpass-python](https://github.com/clarkd/fordpass-python)

## Usage

```go
package main

import (
	"github.com/kacpersaw/go-fordpass"
	"log"
	"os"
)

func main() {
	c := fordpass.NewClient(
		os.Getenv("FORDPASS_USERNAME"),
		os.Getenv("FORDPASS_PASSWORD"),
		os.Getenv("FORDPASS_VIN"))

	status, err := c.Status()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(status.Odometer.Value)
}
```