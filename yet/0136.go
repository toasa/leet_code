func singleNumber(nums []int) int {
	num := nums[0]
	for i := 0; i < len(nums); i++ {
		num ^= nums[i]
	}
	return num\
}