package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/snsinfu/rebed/rebin"
)

type opMode int

const (
	opModeSum opMode = iota
	opModeMean
)

func summarize(beg, end int64, value, weight float64, mode opMode) float64 {
	switch mode {
	case opModeSum:
		return value / float64(end-beg)

	case opModeMean:
		return value / weight
	}
	panic("unexpected mode")
}

func rebed(input io.Reader, output io.Writer, binSize int64, mode opMode) error {
	track, err := rebin.NewTrack(binSize)
	if err != nil {
		return err
	}

	chrom := ""

	printBin := func(beg, end int64, value, weight float64) {
		out := summarize(beg, end, value, weight, mode)
		fmt.Fprintf(output, "%s\t%d\t%d\t%g\n", chrom, beg, end, out)
	}

	s := bufio.NewScanner(input)

	for s.Scan() {
		fields := strings.Fields(s.Text())

		if len(fields) != 4 {
			return errors.New("bad line format")
		}

		if fields[0] != chrom {
			track.GetBins(printBin)
			track, _ = rebin.NewTrack(binSize)
			chrom = fields[0]
		}

		beg, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return err
		}

		end, err := strconv.ParseInt(fields[2], 10, 64)
		if err != nil {
			return err
		}

		val, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return err
		}

		track.Put(beg, end, val)
	}

	if err := s.Err(); err != nil {
		return err
	}

	track.GetBins(printBin)

	return nil
}
