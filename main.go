package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

const version = "0.0.1"

var revision = "HEAD"

func main() {
	cmd := &cli.Command{
		Name:    "hello",
		Usage:   "Hello, World!",
		Version: version,
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
