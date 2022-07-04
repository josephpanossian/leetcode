func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    fmt.Println(nums)
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        target := -1*nums[i]
        head, tail := i+1, len(nums)-1

        for head < tail {
            sum := nums[head] + nums[tail]
            if(sum == target) {
                res = append(res, []int{nums[i], nums[head], nums[tail]})
                head++
                tail--
                for head < tail && nums[head-1] == nums[head] {
                    head++
                }
                for head < tail && nums[tail+1] == nums[tail] {
                    tail--
                }
            } else if sum > target {
                tail--
            } else {
                head++
            }
        }
    }
    return res
}

