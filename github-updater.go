package main

import (
	"github.com/op/go-logging"
)

const (
	version = "0.0.1"
)

var (
	log = logging.MustGetLogger("solutionist")
	args CmdlineArgs
)

func main() {
	args = parseCmdline()
	setupLogging()
	showInfo()
	checkEnvironment()
	downloadGradleBuildTemplate()
	setupDefaultGradleConfig()
	collectGradleConfig()
	patchGradleConfig()
	executeCmd("gradle", `-p`+args.dir+``, "wrapper")
	executeCmd("gradle", `-p`+args.dir+``, "init")
	executeCmd("hg", "init", ``+args.dir+``)
	executeCmd("hg", "addremove", ``+args.dir+``)
	executeCmd("hg", "commit", `-m Start a new Gradle project`, ``+args.dir+``)
	setupDefaultHelgaConfig()
	collectHelgaConfig()
	createHelgaRepo()
}

func showInfo() {
	log.Info(`
            ,    _
           /|   | |
         _/_\_  >_<
        .-\-/.   |
       /  | | \_ |
       \ \| |\__(/
       /('---')  |
      / /     \  |
   _.'  \'-'  /  |
   \----/\=-=/   ' Solutionist ` + version)
	log.Info("Github Updater========================================")
	log.Debug("")
	log.Debug("Using these parameters (use -h for help):")
	log.Debug("%s", args)
}
