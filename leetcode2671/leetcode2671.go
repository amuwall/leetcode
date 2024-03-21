package leetcode2671

type FrequencyTracker struct {
	Numbers     map[int]int
	Frequencies map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		Numbers:     map[int]int{},
		Frequencies: map[int]int{},
	}
}

func (this *FrequencyTracker) Add(number int) {
	if _, ok := this.Numbers[number]; !ok {
		this.Numbers[number] = 0
	}

	oldFrequency := this.Numbers[number]
	this.Numbers[number] += 1
	newFrequency := this.Numbers[number]

	if _, ok := this.Frequencies[oldFrequency]; ok {
		this.Frequencies[oldFrequency]--
		if this.Frequencies[oldFrequency] == 0 {
			delete(this.Frequencies, oldFrequency)
		}
	}

	if _, ok := this.Frequencies[newFrequency]; ok {
		this.Frequencies[newFrequency]++
	} else {
		this.Frequencies[newFrequency] = 1
	}
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if _, ok := this.Numbers[number]; !ok {
		return
	}

	oldFrequency := this.Numbers[number]
	this.Numbers[number] -= 1
	newFrequency := this.Numbers[number]

	if this.Numbers[number] == 0 {
		delete(this.Numbers, number)
	}

	if _, ok := this.Frequencies[oldFrequency]; ok {
		this.Frequencies[oldFrequency]--
		if this.Frequencies[oldFrequency] == 0 {
			delete(this.Frequencies, oldFrequency)
		}
	}
	if newFrequency != 0 {
		if _, ok := this.Frequencies[newFrequency]; ok {
			this.Frequencies[newFrequency]++
		} else {
			this.Frequencies[newFrequency] = 1
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	_, ok := this.Frequencies[frequency]
	return ok
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
