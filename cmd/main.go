package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/snsinfu/rebed"
)

const usage = `
Rebin bedgraph data into uniform bins.

usage:
  rebed [options] [<input> ...]

options:
  --binsize, -b <binsize>   Set bin size. [default: 100]
  --chromsize, -c <sizedb>  Load chrom sizes from file.
  --help, -h                Show this help message and exit.
`

type config struct {
	Inputs  []string `docopt:"<input>"`
	BinSize int64    `docopt:"--binsize"`
	SizeDB  string   `docopt:"--chromsize"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		panic(err)
	}

	var c config
	if err := opts.Bind(&c); err != nil {
		return err
	}

	return rebed.Run(c.Inputs, c.BinSize)
}
