package Week02

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type child struct {
	count      float64
	leftIndex  float64
	rightIndex float64
}

//SubdomainVisits 811. 子域名访问计数
func SubdomainVisits(cpdomains []string) []string {
	var (
		hashMap = make(map[string]int)
		count   int
	)

	for i := 0; i < len(cpdomains); i++ {
		url := ""
		a := strings.Split(cpdomains[i], " ")
		count, _ = strconv.Atoi(a[0])
		tmp := strings.Split(a[1], ".")
		for j := len(tmp) - 1; j >= 0; j-- {
			if j == len(tmp)-1 {
				url += tmp[j]
			} else {
				url = fmt.Sprintf("%s.%s", tmp[j], url)
			}

			hashMap[url] += count
		}
	}

	anw := make([]string, 0, len(hashMap))

	for key, val := range hashMap {
		anw = append(anw, fmt.Sprintf("%d %s", val, key))
	}

	return anw
}

//FindShortestSubArray 697. 数组的度
func FindShortestSubArray(nums []int) int {
	var (
		hashMap    = make(map[int]*child, len(nums))
		maxCount   float64
		ans, index float64
	)

	for i := 0; i < len(nums); i++ {
		index = float64(i)
		if _, has := hashMap[nums[i]]; has {
			hashMap[nums[i]].count++
			hashMap[nums[i]].rightIndex = index
			continue
		}

		hashMap[nums[i]] = &child{1, index, index}
	}

	for _, child := range hashMap {
		if child.count > maxCount {
			maxCount, ans = child.count, child.rightIndex-child.leftIndex+1
		} else if child.count == maxCount {
			ans = math.Min(ans, child.rightIndex-child.leftIndex+1)
		}
	}

	return int(ans)
}

//SubarraySum 560. 和为 K 的子数组
func SubarraySum(nums []int, k int) int {
	var (
		hashMap  = make(map[int]int, len(nums)+1)
		sum, ans int
	)

	hashMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans += hashMap[sum-k]
		hashMap[sum]++
	}

	return ans
}
