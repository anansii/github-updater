package main

import (
	"github.com/op/go-logging"
	"github.com/kardianos/osext"
	"github.com/franela/goreq"
)

const (
// github
	version = "0.0.1"
// githubrepo = "http://github.com/anansii/github-updater"
	githubrepo = "http://api.github.com/repos/topdeskde/releases/latest"
)

var (
	log = logging.MustGetLogger("github-updater")
	args CmdlineArgs
	exePath string
	dirPath string
)

func main() {
	args = parseCmdline()
	setupLogging()
	showInfo()
	findPath()
	// download latest release to this exe folder
	// rename this executable
	// rename new executable
}

func showInfo() {
	log.Info("Github Updater v%s", version)
	log.Debug("")
	log.Debug("Using these parameters (use -h for help):")
	log.Debug("%s", args)
}

func findPath() {
	path, err := osext.Executable()
	if err != nil {
		log.Fatalf("Failed to find myself: %s", err)
	}
	log.Debug("This executable: %v", path)
	exePath = path

	path, err = osext.ExecutableFolder()
	if err != nil {
		log.Fatalf("Failed to find my directory: %s", err)
	}
	log.Debug("This directory: %v", path)
	dirPath = path
}

func latestReleaseData() {
	res, err := goreq.Request{Uri: githubrepo }.Do()
}

