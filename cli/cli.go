package cli

import (
	"flag"
)

type Args struct {
	Update      bool
	ShowVersion bool
}

func Parse() *Args {
	var args Args

	flag.BoolVar(&args.Update, "u", false, "Check and apply update")
	flag.BoolVar(&args.Update, "update", false, "Check and apply update")
	flag.BoolVar(&args.ShowVersion, "v", false, "Show version")
	flag.BoolVar(&args.ShowVersion, "version", false, "Show version")
	flag.Parse()

	return &args
}
