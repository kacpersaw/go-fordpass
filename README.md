# go-fordpass
A Go client library for the FordPass API. This library provides methods to return vehicle status, start/stop engine and lock/unlock doors.

Tested with `Ford Focus MK4`

Project based on [fordpass-python](https://github.com/clarkd/fordpass-python)

## Features
    * Get vehicle status
    * Start/Stop engine (if supported)
    * Lock/Unlock doors

## Usage

```go
package main

import (
	"github.com/kacpersaw/go-fordpass/fordpass"
	"log"
	"os"
)

func main() {
	car := fordpass.NewClient(
		os.Getenv("FORDPASS_USERNAME"),
		os.Getenv("FORDPASS_PASSWORD"),
		os.Getenv("FORDPASS_VIN"))

	status, err := car.Status()
	if err != nil {
		log.Fatal(err)
	}

	//Get odometer value
	log.Println(status.Odometer.Value)

	//Start engine
	err = car.Start()
	if err != nil {
		log.Fatal(err)
	}
	
	//Stop engine
	err = car.Stop()
	if err != nil {
		log.Fatal(err)
	}
	
	//Lock doors
	err = car.Lock()
	if err != nil {
		log.Fatal(err)
	}
	
	//Unlock doors
	err = car.Unlock()
	if err != nil {
		log.Fatal(err)
	}
}
```