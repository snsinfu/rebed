package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/pkg/errors"
)

const usage = `
Rebin bedgraph into uniform bins.

Usage:
  rebed [options] [<input>...]

Options:
  -b, --binsize <binsize>  Set bin size. [default: 1000]
  -m, --mode <mode>        Set binning mode (mean or sum). [default: sum]
  -h, --help               Show this help message and exit.
`

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

	// Bin size
	binSize, err := opts.Int("--binsize")
	if err != nil {
		return fmt.Errorf("bad bin size")
	}

	if binSize <= 0 {
		return fmt.Errorf("bad bin size %d", binSize)
	}

	// Binning mode
	modeName, err := opts.String("--mode")
	if err != nil {
		return err
	}

	var mode opMode
	switch modeName {
	case "sum":
		mode = opModeSum

	case "mean":
		mode = opModeMean

	default:
		return fmt.Errorf("bad mode: %s", modeName)
	}

	// Inputs
	inputs := opts["<input>"].([]string)

	if len(inputs) == 0 {
		if err := rebed(os.Stdin, os.Stdout, int64(binSize), mode); err != nil {
			return errors.Wrap(err, "processing stdin")
		}
	}

	for _, input := range inputs {
		err := func() error {
			file, err := os.Open(input)
			if err != nil {
				return err
			}
			defer file.Close()

			return rebed(file, os.Stdout, int64(binSize), mode)
		}()
		if err != nil {
			return errors.Wrap(err, "processing file "+input)
		}
	}

	return nil
}
