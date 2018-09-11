package rebed

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func Run(inputs []string, binSize int64) error {
	for _, input := range inputs {
		if err := rebed(input, binSize); err != nil {
			return err
		}
	}
	return nil
}

func rebed(input string, binSize int64) error {
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	curChrom := ""
	chromSize := int64(0)
	bins := []float64{}

	for s.Scan() {
		chrom, begin, end, value, err := parseLine(s.Text())
		if err != nil {
			return err
		}

		if chrom != curChrom {
			outputBins(curChrom, bins, binSize, chromSize)

			curChrom = chrom
			chromSize = 0
			bins = bins[:0]
		}

		n := int((end + binSize - 1) / binSize)
		if extend := n - len(bins); extend > 0 {
			bins = append(bins, make([]float64, extend)...)
		}

		curBegin := begin
		curEnd := (begin/binSize + 1) * binSize

		for curBegin < end {
			i := curBegin / binSize
			bins[i] += value * float64(curEnd-curBegin)

			curBegin = curEnd
			curEnd += binSize
			if curEnd > end {
				curEnd = end
			}
		}

		if chromSize < end {
			chromSize = end
		}
	}

	outputBins(curChrom, bins, binSize, chromSize)

	return nil
}

func outputBins(chrom string, bins []float64, binSize int64, chromSize int64) {
	for i, sum := range bins {
		begin := int64(i) * binSize
		end := begin + binSize

		if end > chromSize {
			end = chromSize
		}

		span := end - begin
		mean := sum / float64(span)

		fmt.Printf("%s\t%d\t%d\t%g\n", chrom, begin, end, mean)
	}
}

func parseLine(s string) (chrom string, begin, end int64, value float64, err error) {
	fields := strings.Fields(s)
	if len(fields) != 4 {
		err = errors.New("invalid number of fields")
		return
	}

	chrom = fields[0]

	begin, err = strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return
	}

	end, err = strconv.ParseInt(fields[2], 10, 64)
	if err != nil {
		return
	}

	value, err = strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return
	}

	return
}
