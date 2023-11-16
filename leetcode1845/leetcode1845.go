package leetcode1845

type SeatManager struct {
	availableSeatNumber int
	seats               []bool
}

func Constructor(n int) SeatManager {
	return SeatManager{
		availableSeatNumber: 0,
		seats:               make([]bool, n),
	}
}

func (this *SeatManager) Reserve() int {
	seatNumber := this.availableSeatNumber
	if seatNumber == len(this.seats) {
		return -1
	}
	this.seats[seatNumber] = true

	for {
		this.availableSeatNumber++
		if this.availableSeatNumber == len(this.seats) {
			break
		}
		if !this.seats[this.availableSeatNumber] {
			break
		}
	}

	return seatNumber + 1
}

func (this *SeatManager) Unreserve(seatNumber int) {
	seatNumber = seatNumber - 1
	if !this.seats[seatNumber] {
		return
	}

	this.seats[seatNumber] = false
	if seatNumber < this.availableSeatNumber {
		this.availableSeatNumber = seatNumber
	}
}
