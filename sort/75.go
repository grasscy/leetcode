package sort

func sortColors(nums []int) []int {
	i, j := 0, len(nums)-1

	zero := 0

	for i <= j {
		if nums[i] == 0 {
			nums[zero], nums[i] = nums[i], nums[zero]
			zero++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}

	}

	return nums
}
