package main

import (
	"fmt"
	"flag"
	"os"
)

func Usage() {
	fmt.Println("Usage: unpackage CMD [ARGS]...")
}

func main() {
	flag.Usage = Usage

	args := flag.Args()
	if len(args) == 0 {
		Usage()
		os.Exit(1)
	}

	switch args[1] {
	case "add":
		fallthrough
	case "load":
		environment.LoadModule(args)

	case "reload":
		environment.ReloadModule(args)

	case "remove":
		fallthrough
	case "rm":
		fallthrough
	case "unload":
		environment.UnloadModule(args)

	case "purge":
		fallthrough
	case "reset":
		environment.UnloadAllModules()

	case "info":
		fallthrough
	case "show":
		environment.ShowModule(args)

	case "list":
		environment.ListModules()

	case "avail":
		environment.AvailableModules(args)

	default:
		Usage()
		os.Exit(1)
	}
	os.Exit(0)
}