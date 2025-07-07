package main

func MostFrequentNaive(nums []int) int {
    var keyCount [2]int
    for left, currentNum := range nums {
        count := 1
        for _, rightNum := range nums[left+1:] {
            if currentNum == rightNum {
                count++
            }
        }
        if count > keyCount[1] {
            keyCount = [2]int{currentNum, count}
        }
    }
    return keyCount[0]
}

func MostFrequent(nums []int) int {
    keyCountMap := make(map[int]int)
    var keyCountMax [2]int
    for _, num := range nums {
        keyCountMap[num]++
    }
    for key, val := range keyCountMap {
        if val > keyCountMax[1] {
            keyCountMax = [2]int{key, val}
        }
    }
    return keyCountMax[0]
}
