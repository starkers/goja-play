package main

import (
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/ghodss/yaml"
	flag "github.com/spf13/pflag"
	"os"
)

var flagFile string

func init() {
	flag.StringVar(&flagFile, "file", "", "which file to run")
}

var log = StartLogger(true)

func main() {
	flag.Parse()
	jsCompileStrict := true
	if flagFile == "" {
		log.Fatal("please specify --file foo.js")
	}

	rawFile, err := os.ReadFile(flagFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fileContentsString := string(rawFile)

	vm := goja.New()
	reg := require.NewRegistry(
	// require.WithGlobalFolders("."),
	)
	reg.Enable(vm)
	console.Enable(vm) // support "console()" inside js

	compiledProgram, err := goja.Compile("", fileContentsString, jsCompileStrict)
	if err != nil {
		log.Errorw("Error compiling", "file", flagFile, "err", err)
		return
	}
	result, err := vm.RunProgram(compiledProgram)
	log.Debugw("vm.RunProgram",
		"file", flagFile,
		"result", result,
		"err", err,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Pulls function "getData" out
	getData, ok := goja.AssertFunction(vm.Get("getData"))
	if !ok {
		log.Fatalw("couldn't find function 'getData()' in the javascript",
			"file", flagFile)
	}

	res, err := getData(goja.Undefined())
	if err != nil {
		log.Fatalw("getData(): Unable to call getData() in js",
			"file", flagFile,
			"err", err)
	}

	log.Infow("got the following resources",
		"file", flagFile,
		"res", res)
	prettyPrintJSONString(res.ToString().String())
}

func prettyPrintJSONString(input string) {
	fmt.Println("--- json retrieved from js follows ---")
	var prettyObj map[string]interface{}
	err := yaml.Unmarshal([]byte(input), &prettyObj)
	if err != nil {
		log.Errorw("Couldn't decode yaml or json from input string",
			"string", input,
		)
		return
	}
	f := colorjson.NewFormatter()
	f.Indent = 2

	s, _ := f.Marshal(prettyObj)
	fmt.Println(string(s))
	fmt.Println("---")
}
