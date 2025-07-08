package arrays

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
    var maxKey, count int
    for _, num := range nums {
        keyCountMap[num]++
        if keyCountMap[num] > count {
            maxKey, count = num, keyCountMap[num]
        }
    }
    return maxKey
}
