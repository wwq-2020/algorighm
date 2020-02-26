package subsets

func subsets(nums []int) [][]int {
	var track []int
	var res [][]int
	backTrack(nums, 0, track, &res)
	return res
}

func backTrack(nums []int, start int, track []int, res *[][]int) {
	*res = append(*res, track)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		backTrack(nums, i+1, track, res)
		track = (track)[:len(track)-1]
	}
}
