package main

import (
	"fmt"
	"log"

	"go_updates/cli"
	"go_updates/updater"
)

var version = "0.0.1-dev" // Set via -ldflags="-X main.version=..."

func main() {
	args := cli.Parse()

	if args.ShowVersion {
		fmt.Println(version)
		return
	}

	if args.Update {
		if err := updater.Update(version); err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("Pretending to do some stuff...")
}
