package rebin

import (
	"reflect"
	"testing"
)

func TestNewTrack_createsEmptyTrack(t *testing.T) {
	track, err := NewTrack(1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	track.GetBins(func(beg, end int64, val float64) {
		t.Errorf("unexpected bin: (%d, %d, %g)", beg, end, val)
	})
}

func TestNewTrack_detectsBadBinSize(t *testing.T) {
	// Zero
	if track, err := NewTrack(0); err == nil {
		t.Errorf("unexpected success: track = %v", track)
	}

	// Negative
	if track, err := NewTrack(-1); err == nil {
		t.Errorf("unexpected success: track = %v", track)
	}
}

func TestTrack_Put_GetBins(t *testing.T) {
	track, err := NewTrack(5)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	type bin struct {
		beg int64
		end int64
		val float64
	}

	input := []bin{
		{4, 9, 1.0},
		{0, 3, 2.0},
		{3, 9, 3.0},
		{2, 6, 4.0},
		{9, 12, 5.0},
	}

	expected := []bin{
		{0, 5, 5.0},
		{5, 10, 5.0},
		{10, 12, 5.0},
	}

	for _, b := range input {
		if err := track.Put(b.beg, b.end, b.val); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	actual := []bin{}

	track.GetBins(func(beg, end int64, val float64) {
		actual = append(actual, bin{
			beg: beg,
			end: end,
			val: val,
		})
	})

	if len(actual) != len(expected) {
		t.Fatalf("unexpected number of bins: got %d, want %d", len(actual), len(expected))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected bins: got %v, want %v", actual, expected)
	}
}

func TestTrack_Put_detectsBadInterval(t *testing.T) {
	track, err := NewTrack(10)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Negative coord
	if err := track.Put(-1, 1, 1.0); err == nil {
		t.Fatalf("unexpected success: track = %v", track)
	}

	// Reversed order
	if err := track.Put(2, 1, 1.0); err == nil {
		t.Fatalf("unexpected success: track = %v", track)
	}
}
