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
	bins    []Bin
	binSize int64
	end     int64
}

type Bin struct {
	value  float64
	weight float64
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
func (t *Track) GetBins(callback func(int64, int64, float64, float64)) {
	for beg := int64(0); beg < t.end; beg += t.binSize {
		end := beg + t.binSize
		if end > t.end {
			end = t.end
		}

		i := int(beg / t.binSize)
		callback(beg, end, t.bins[i].value, t.bins[i].weight)
	}
}

// Put accumulates value val over half-open interval [beg, end). It fails if
// beg < 0 or beg > end.
func (t *Track) Put(beg, end int64, val float64) error {
	if err := t.Cover(beg, end); err != nil {
		return err
	}

	b := beg
	e := (beg/t.binSize + 1) * t.binSize

	for b < end {
		if e > end {
			e = end
		}

		i := int(b / t.binSize)
		w := float64(e - b)
		t.bins[i].value += val * w
		t.bins[i].weight += w

		b = e
		e += t.binSize
	}

	return nil
}

// Cover ensures specified coordinate range [beg, end) to be covered by the
// track by extending the end of the track if necessary. If you do not use
// this function, Track defines its end as the highest end coordinate of
// samples seen.
func (t *Track) Cover(beg, end int64) error {
	if beg < 0 || beg > end {
		return errBadInterval
	}

	bound := int((end + t.binSize - 1) / t.binSize)
	if bound > len(t.bins) {
		newBins := make([]Bin, bound+len(t.bins))
		copy(newBins, t.bins)
		t.bins = newBins
	}

	if end > t.end {
		t.end = end
	}

	return nil
}
