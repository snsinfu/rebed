package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// loadChromSizes loads mapping of chromosome name to its size from a
// tab-separated text file (in the format of hg19.chrom.sizes).
func loadChromSizes(filename string) (map[string]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	chromSizes := map[string]int64{}

	s := bufio.NewScanner(file)
	for s.Scan() {
		fields := strings.Fields(s.Text())

		if len(fields) != 2 {
			return chromSizes, errors.New("bad line format")
		}

		chrom := fields[0]
		size, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return chromSizes, err
		}

		chromSizes[chrom] = size
	}

	if err := s.Err(); err != nil {
		return chromSizes, err
	}

	return chromSizes, nil
}
