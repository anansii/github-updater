package main

import (
	"github.com/op/go-logging"
	"os"
)

type Hidden string

func (h Hidden) Redacted() interface{} {
	return logging.Redact(string(h))
}

func setupLogging() {

	consoleFormat := logging.MustStringFormatter(" %{level: 8s} | %{message}")
	consoleBackend := logging.NewLogBackend(os.Stdout, "", 0)
	consoleBackendFormatted := logging.NewBackendFormatter(consoleBackend, consoleFormat)
	consoleBackendLeveled := logging.AddModuleLevel(consoleBackendFormatted)
	consoleBackendLeveled.SetLevel(logging.INFO, "")
	if args.debug {
		logging.SetBackend(consoleBackendFormatted)
	} else {
		logging.SetBackend(consoleBackendLeveled)
	}

}
