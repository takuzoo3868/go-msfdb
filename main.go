package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/takuzoo3868/go-msfdb/commands"
)

// Name :
const Name string = "go-msfdb"

var version = "0.0.1"

func main() {
	var v = flag.Bool("v", false, "Show version")

	if envArgs := os.Getenv("GOVAL_DICTIONARY_ARGS"); 0 < len(envArgs) {
		if err := flag.CommandLine.Parse(strings.Fields(envArgs)); err != nil {
			fmt.Printf("Failed to get ENV Vars: %s", err)
			os.Exit(1)
		}
	} else {
		flag.Parse()
	}
	
	if *v {
		fmt.Printf("go-msfdb %s \n", version)
		os.Exit(0)
	}

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
