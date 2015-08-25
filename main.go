package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"text/template"
)

const Version = "0.1.0"
const extension = ".tpl"

func usageAndExit(s string) {
	fmt.Fprintf(os.Stderr, "!! %s\n", s)
	fmt.Fprint(os.Stderr, "usage: envtpl [template]\n")
	fmt.Fprintf(os.Stderr, "%s version: %s (%s on %s/%s; %s)\n", os.Args[0], Version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler)
	os.Exit(1)
}

func init() {
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		usageAndExit("missing [template]")
	}
	input := args[0]

	if input[len(input) - len(extension):] != extension {
		usageAndExit("[template] does not end with " + extension)
	}

	info, err := os.Stat(input)
	if err != nil {
		usageAndExit(err.Error())
	}

	t, err := template.ParseFiles(input)
	if err != nil {
		usageAndExit(err.Error())
	}

	var b bytes.Buffer
	err = t.Option("missingkey=zero").Execute(&b, getEnvironMap())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ioutil.WriteFile(input[:len(input) - len(extension)], b.Bytes(), info.Mode())
	os.Remove(input)
}

func getEnvironMap() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		x := strings.SplitN(e, "=", 2)
		env[x[0]] = x[1]
	}
	return env
}
