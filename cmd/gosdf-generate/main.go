package main

import (
	"github.com/Jozmen/gosdf/pkg/cmd/generate"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := generate.GenerateCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
