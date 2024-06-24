package main

import (
	"flag"
	"os"
)

var CommandLine = NewFlagSet(os.Args[0], flag.ExitOnError)

var port int = 8080

const host string = "127.0.0.1"

func init() {

}

func NewFlagSet(name string, handling flag.ErrorHandling) *flag.FlagSet {
	return nil
}
