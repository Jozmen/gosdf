package main

import (
	"github.com/Jozmen/gosdf/pkg/cmd/convert"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := convert.RootCommand.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
