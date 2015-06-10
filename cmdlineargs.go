package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"strings"
)

type CmdlineArgs struct {
	fakeversion string
	selfupdate  bool
	debug       bool

}

func (a CmdlineArgs) String() string {
	return fmt.Sprintf("Commandline arguments: %+v", a)
}

func parseCmdline() CmdlineArgs {
	fakeversion := flag.String("fakeversion", version, "Version string used for update check")
	githubrepo := flag.String("githubrepo", githubrepo, "Github repository used to update")
	selfupdate := flag.Bool("selfupdate", false, "Update this executable from Github")
	debug := flag.Bool("debug", false, "Show debug information")
	flag.Parse()

	return CmdlineArgs{fakeversion: *fakeversion, githubrepo: *githubrepo, selfupdate: *selfupdate, debug: *debug}
}
