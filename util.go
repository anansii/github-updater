package main

import (
	"bufio"
	"fmt"
	"github.com/bgentry/speakeasy"
	"github.com/franela/goreq"
	"io"
	"os"
	"os/exec"
	"strings"
)

func requestInput(value *string, description string) {
	log.Warning(description)
	log.Info("[%v]", *value)
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	if err != nil {
		log.Critical("Error: %v", err)
		log.Fatal("No reason to go on. This ends now :(")
	}
	log.Debug("Value provided: %v", string(input))
	log.Debug("Value provided: %v", input)
	log.Debug("Value length: %d", len(input))
	if len(input) != 0 {
		*value = string(input)
	}
}

func requestHiddenInput(value *string, description string) {
	log.Warning(description)
	input, err := speakeasy.Ask("> ")
	if err != nil {
		log.Critical("Error: %v", err)
		log.Fatal("No reason to go on. This ends now :(")
	}
	log.Debug("Value provided: %v", Hidden(input))
	log.Debug("Value length: %d", len(input))
	if len(input) != 0 {
		*value = input
	}
}

// general make http request

func executeCmd(cmdName string, cmdArgs ...string) {
	cmd := exec.Command(cmdName, cmdArgs...)

	log.Notice("> Executing: %s", strings.Join(cmd.Args, " "))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("%s\nThis ended abruptly.", err)
	}
}

func downloadFromUrl(url string, targetDir string, fileName string, username string, password string) {
	if targetDir == "" {
		targetDir = "."
	}
	if fileName == "" {
		tokens := strings.Split(url, "/")
		fileName = tokens[len(tokens)-1]
	}

	targetPath := targetDir + "/" + fileName

	log.Debug("Downloading from [%s] to [%s] using [%s:%s]", url, targetPath, username, Hidden(password))

	res, err := goreq.Request{
		Method:            "GET",
		Uri:               url,
		BasicAuthUsername: username,
		BasicAuthPassword: password,
	}.Do()
	defer res.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	if res.StatusCode == 200 {
		file, err := os.Create(targetPath)
		if err != nil {
			log.Panicf("Failed to create %s: %s", targetPath, err)
		}
		defer file.Close()

		size, err := io.Copy(file, res.Body)
		if err != nil {
			log.Panicf("Failed to write downloaded data to %s: %s", targetPath, err)
		}

		log.Notice("%s with %v bytes downloaded", targetPath, size)
	} else {
		log.Panicf("Error: %s", res.Status)
	}

}

/*
func downloadGradleBuildTemplate() {

	url := "http://helga/scm/hg/gradle/solution-plugin/raw-file/tip/setup/template-build.gradle"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(args.username, args.password)
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		file, err := os.Create(args.dir + "/build.gradle")
		if err != nil {
			log.Panic("Failed to create build.gradle: %s", err)
		}
		defer file.Close()

		size, err := io.Copy(file, resp.Body)
		if err != nil {
			log.Panic("Failed to write downloaded data to build.gradle: %s", err)
		}

		log.Notice("build.gradle with %v bytes downloaded", size)
	} else {
		log.Critical("Error: %s", resp.Status)
		log.Fatal("No reason to go on. This ends now :(")
	}
}
*/

/*
func handle(e error) {
    if e != nil {
        panic(e)
    }
}
*/
