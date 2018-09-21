package rebin

import (
	"errors"
)

var (
	errBadBinSize  = errors.New("bad bin size")
	errBadInterval = errors.New("bad interval")
)

// Track
type Track struct {
	bins    []float64
	binSize int64
	end     int64
}

// NewTrack creates an empty Track with given bin size. It fails if binSize is
// not positive.
func NewTrack(binSize int64) (*Track, error) {
	if binSize <= 0 {
		return nil, errBadBinSize
	}
	return &Track{binSize: binSize}, nil
}

// GetBins calls callback for all bins.
func (t *Track) GetBins(callback func(int64, int64, float64)) {
	for beg := int64(0); beg < t.end; beg += t.binSize {
		end := beg + t.binSize
		if end > t.end {
			end = t.end
		}

		i := int(beg / t.binSize)
		mean := t.bins[i] / float64(end-beg)

		callback(beg, end, mean)
	}
}

// Put accumulates value val over half-open interval [beg, end). It fails if
// beg < 0 or beg > end.
func (t *Track) Put(beg, end int64, val float64) error {
	if beg < 0 || beg > end {
		return errBadInterval
	}

	bound := int((end + t.binSize - 1) / t.binSize)
	if bound > len(t.bins) {
		newBins := make([]float64, bound+len(t.bins))
		copy(newBins, t.bins)
		t.bins = newBins
	}

	b := beg
	e := (beg/t.binSize + 1) * t.binSize

	for b < end {
		if e > end {
			e = end
		}

		i := int(b / t.binSize)
		t.bins[i] += val * float64(e-b)

		b = e
		e += t.binSize
	}

	if end > t.end {
		t.end = end
	}

	return nil
}
