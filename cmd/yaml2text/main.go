package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mashiike/yaml2text"
)

func main() {
	var (
		templateFile string
	)
	f := newFlagSet()
	f.StringVar(&templateFile, "template", "default.tpl", "go template file for yaml convert")
	f.Parse(os.Args[1:])
	if f.NArg() < 1 {
		f.Usage()
		return
	}

	app, err := yaml2text.NewWithFile(templateFile)
	if err != nil {
		f.Usage()
		log.Fatal(err)
	}
	if err := app.ExecuteWithFile(f.Arg(0), os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func newFlagSet() *flag.FlagSet {
	f := flag.NewFlagSet("yaml2text", flag.ExitOnError)
	f.Usage = func() {
		fmt.Fprintf(f.Output(), "Usage: %s [-template <path>] <yaml path>\n\n", f.Name())
		fmt.Fprintln(f.Output(), "This is a yaml convertor with go template format.")
		fmt.Fprintln(f.Output(), "")
		f.PrintDefaults()
	}
	return f
}
