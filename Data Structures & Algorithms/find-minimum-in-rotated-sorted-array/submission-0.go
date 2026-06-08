func findMin(nums []int) int {
    l := 0
    r := len(nums) - 1

    for l < r {
        mid := l + (r-l)/2

        if nums[mid] > nums[len(nums)-1] {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return nums[l]
}