package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
)

type CmdlineArgs struct {
	fakeVersion string
	selfupdate  bool
	debug       bool

}

func (a CmdlineArgs) String() string {
	args := fmt.Sprintf("dir=%s\n", a.dir)
	args += fmt.Sprintf("username=%s\n", a.username)
	args += fmt.Sprintf("password=%s\n", a.maskedPassword())
	args += fmt.Sprintf("logfile=%v\n", a.logfile)
	args += fmt.Sprintf("debug=%v\n", a.debug)
	args += fmt.Sprintf("color=%v\n", a.color)
	return args
}

func (a CmdlineArgs) maskedPassword() string {
	return strings.Repeat("*", len(a.password))
}

func parseCmdline() CmdlineArgs {
	defaultUsername := ""
	currentUser, err := user.Current()
	if err == nil {
		usernameParts := strings.Split(currentUser.Username, "\\")
		defaultUsername = strings.ToLower(usernameParts[len(usernameParts)-1])
	}

	dir := flag.String("dir", ".", "Target directory to create project in; defaults to current directory")
	username := flag.String("username", defaultUsername, "Username used for authentication")
	password := flag.String("password", "", "Password used for authentication")
	logfile := flag.Bool("logfile", false, "Logs output to logile in project directory")
	debug := flag.Bool("debug", false, "Show debug information")
	color := flag.Bool("color", false, "Use colors in output. Uses ANSI escape sequences")
	flag.Parse()

	err = os.MkdirAll(*dir, 0777)
	if err != nil {
		log.Fatal("Target directory could not be created: %s", err)
	}

	return CmdlineArgs{dir: *dir, username: *username, password: *password, logfile: *logfile, debug: *debug, color: *color}
}
