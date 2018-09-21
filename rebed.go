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

func rebed(input io.Reader, output io.Writer, binSize int64) error {
	track, err := rebin.NewTrack(binSize)
	if err != nil {
		return err
	}

	chrom := ""

	printBin := func(beg, end int64, val float64) {
		fmt.Fprintf(output, "%s\t%d\t%d\t%g\n", chrom, beg, end, val)
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
