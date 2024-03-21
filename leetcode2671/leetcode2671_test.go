package leetcode2671

import (
	"testing"
)

func TestFrequencyTracker(t *testing.T) {
	// example 1
	tracker := Constructor()
	tracker.Add(3)
	tracker.Add(3)
	if got := tracker.HasFrequency(2); got != true {
		t.Errorf("HasFrequency() = %v, want %v", got, true)
		return
	}

	// example 2
	tracker = Constructor()
	tracker.Add(1)
	tracker.DeleteOne(1)
	if got := tracker.HasFrequency(1); got != false {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}

	// example 3
	tracker = Constructor()
	if got := tracker.HasFrequency(2); got != false {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
	tracker.Add(3)
	if got := tracker.HasFrequency(1); got != true {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}

	// example 4
	tracker = Constructor()
	if got := tracker.HasFrequency(1); got != false {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
	tracker.Add(3)
	if got := tracker.HasFrequency(1); got != true {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
	if got := tracker.HasFrequency(1); got != true {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
	tracker.Add(6)
	tracker.Add(2)
	tracker.Add(6)
	tracker.DeleteOne(6)
	tracker.DeleteOne(6)
	if got := tracker.HasFrequency(2); got != false {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
	tracker.Add(2)
	if got := tracker.HasFrequency(2); got != true {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
	if got := tracker.HasFrequency(1); got != true {
		t.Errorf("HasFrequency() = %v, want %v", got, false)
		return
	}
}
