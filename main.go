package main

import (
	"flag"
	"io/ioutil"
	"os"

	kronk "github.com/kcmerrill/kronk/core"
)

func main() {
	out := flag.String("out", "simple", "Output format. (simple|csv|tsv|inline)")
	del := flag.String("del", ",", "Delimiter to use for output")
	passThru := flag.Bool("pass-through", false, "If there are errors, pass through results")
	flag.Parse()

	if *out == "simple" && len(flag.Args()) >= 2 {
		// saving you from yourself yo
		*out = "tsv"
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// stdin, we should use it!
		stdin, _ := ioutil.ReadAll(os.Stdin)
		kronk.NewKronk(flag.Args(), stdin).Display(*out, *del, *passThru)
		return
	}
	kronk.Says("error", "Unable to parse stdin")
	os.Exit(1)
}
