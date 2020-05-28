package main

func smallerNumbersThanCurrent(nums []int) []int {
    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if nums[j] < nums[i]{
                result[i] += 1                
            }
        }
    }
    return result
}


// Time Complexity: O(NlogN) (N + NlogN + N + N)
// Space Complexity: O(N) (N + logN + N + answer)
func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	countMap := map[int]int{}
	for i, num := range sortedNums {
	  if _, ok := countMap[num]; !ok {
		countMap[num] = i
	  }
	}
	counts := make([]int, len(nums))
	for i, num := range nums {
	  counts[i] = countMap[num]
	}
	return counts
}

// Time Complexity: O(N) (N + 100 + N)
// Space Complexity: O(1) (101 + answer)
func smallerNumbersThanCurrent(nums []int) []int {
	occurrences := [101]int{}
	for _, num := range nums {
	  occurrences[num]++
	}
	prevOccurrences := occurrences[0]
	occurrences[0] = 0
	for i := range occurrences[1:] {
	  occurrences[i+1], prevOccurrences = prevOccurrences, prevOccurrences+occurrences[i+1]
	}
	counts := make([]int, len(nums))
	for i, num := range nums {
	  counts[i] = occurrences[num]
	}
	return counts
}