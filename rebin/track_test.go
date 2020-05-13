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

	track.GetBins(func(beg, end int64, v float64, w float64) {
		t.Errorf("unexpected bin: (%d, %d, %g, %g)", beg, end, v, w)
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

	type inputBin struct {
		beg int64
		end int64
		val float64
	}

	type outputBin struct {
		beg    int64
		end    int64
		val    float64
		weight float64
	}

	//              1    1
	//    0    5    0    5
	//    |----|----|----|
	// I      ===== 1.0
	// N  === 2.0
	// P     ====== 3.0
	// U    ==== 4.0
	// T           === 5.0
	//    |----|----|----|
	// v  | 25 | 25 | 10 |
	// w  |  9 | 10 |  2 |

	input := []inputBin{
		{4, 9, 1.0},
		{0, 3, 2.0},
		{3, 9, 3.0},
		{2, 6, 4.0},
		{9, 12, 5.0},
	}

	expected := []outputBin{
		{0, 5, 25.0, 9.0},
		{5, 10, 25.0, 10.0},
		{10, 12, 10.0, 2.0},
	}

	for _, b := range input {
		if err := track.Put(b.beg, b.end, b.val); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	actual := []outputBin{}

	track.GetBins(func(beg, end int64, v float64, w float64) {
		actual = append(actual, outputBin{
			beg:    beg,
			end:    end,
			val:    v,
			weight: w,
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

func TestTrack_Cover_extendsTrack(t *testing.T) {
	track, err := NewTrack(10)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	coverStart := int64(5)
	coverEnd := int64(23)
	if err := track.Cover(coverStart, coverEnd); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	trackEnd := int64(-1)
	track.GetBins(func(beg, end int64, v float64, w float64) {
		trackEnd = end
	})

	if trackEnd != coverEnd {
		t.Errorf("unexpected track end: got %v, want %v", trackEnd, coverEnd)
	}
}

func TestTrack_Cover_detectsBadInterval(t *testing.T) {
	track, err := NewTrack(10)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Negative coord
	if err := track.Cover(-1, 1); err == nil {
		t.Fatalf("unexpected success: track = %v", track)
	}

	// Reversed order
	if err := track.Cover(2, 1); err == nil {
		t.Fatalf("unexpected success: track = %v", track)
	}
}
