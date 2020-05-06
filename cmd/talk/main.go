package main

import (
	"github.com/devplayg/hippo"
	"github.com/devplayg/messenger"
)

func main() {
	hub := messenger.NewHub()
	config := &hippo.Config{
		Name:        "t",
		Description: "tt",
		Version:     "beta",
		//LogDir:      "",
		Debug: true,
		//Trace:       false,
		//Logger:      nil,
		//CertFile:    "",
		//KeyFile:     "",
		//Insecure:    true,
	}
	engine := hippo.NewEngine(hub, config)
	if err := engine.Start(); err != nil {
		hub.Log.Error(err)
	}
}
