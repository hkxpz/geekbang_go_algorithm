package Week05

import "sort"

//1011. 在 D 天内送达包裹的能力
func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	return left + sort.Search(right-left, func(x int) bool {
		x += left
		day := 1
		sum := 0
		for _, w := range weights {
			if sum+w > x {
				day++
				sum = 0
			}
			sum += w
		}
		return day <= days
	})
}

//911. 在线选举
type TopVotedCandidate struct {
	persons []int
	times   []int
	result  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	sta := map[int]int{}
	result := make([]int, len(times))
	sta[persons[0]] = 1
	result[0] = persons[0]
	for i := 1; i < len(persons); i++ {
		lastTimes, _ := sta[result[i-1]]
		maxTimes := lastTimes

		times, _ := sta[persons[i]]
		if times+1 >= maxTimes {
			result[i] = persons[i]
		} else {
			result[i] = result[i-1]
		}

		sta[persons[i]] += 1

	}

	return TopVotedCandidate{
		persons: persons,
		times:   times,
		result:  result,
	}

}

func (c *TopVotedCandidate) Q(t int) int {
	index := sort.SearchInts(c.times, t)
	if index == len(c.result) {
		index--
	} else {
		if c.times[index] != t {
			index--
		}
	}

	return c.result[index]
}

//875. 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	ceil := func(x, y int) int {
		if x%y == 0 {
			return x / y
		}
		return x/y + 1
	}
	check := func(x int) bool {
		if x == 0 {
			return false
		}
		t := 0
		for i := 0; i < n; i++ {
			t += ceil(piles[i], x)
		}
		return t <= h
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += piles[i]
	}
	l, r := sum/h, sum
	for l < r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if check(l) {
		return l
	}
	return -1
}
