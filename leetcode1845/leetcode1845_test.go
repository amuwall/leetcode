package leetcode1845

import (
	"testing"
)

func TestSeatManager(t *testing.T) {
	seatManager := Constructor(5)

	seatNumber := seatManager.Reserve()
	if seatNumber != 1 {
		t.Errorf("Reserve() = %v, want %v", seatNumber, 1)
	}

	seatNumber = seatManager.Reserve()
	if seatNumber != 2 {
		t.Errorf("Reserve() = %v, want %v", seatNumber, 2)
	}

	seatManager.Unreserve(2)

	seatNumber = seatManager.Reserve()
	if seatNumber != 2 {
		t.Errorf("Reserve() = %v, want %v", seatNumber, 2)
	}

	seatNumber = seatManager.Reserve()
	if seatNumber != 3 {
		t.Errorf("Reserve() = %v, want %v", seatNumber, 3)
	}

	seatNumber = seatManager.Reserve()
	if seatNumber != 4 {
		t.Errorf("Reserve() = %v, want %v", seatNumber, 4)
	}

	seatNumber = seatManager.Reserve()
	if seatNumber != 5 {
		t.Errorf("Reserve() = %v, want %v", seatNumber, 5)
	}
}
