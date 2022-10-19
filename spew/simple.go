/*
bash$ A=1 B=True go run simple.go

(main.Env) {
	A: (int) 1,
	B: (bool) true,
	C: (string) (len=2) "80"
	}
*/

package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	// Common settings
	A int    `envconfig:"A" default:"1"`
	B bool   `envconfig:"B" default:"false"`
	C string `envconfig:"C" default:"80"`
}

func main() {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
		return
	}

	scs := spew.ConfigState{DisableCapacities: true, DisableMethods: true}
	scs.Dump(env)
}
