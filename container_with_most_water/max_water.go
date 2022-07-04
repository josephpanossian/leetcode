func maxArea(height []int) int {
    head, tail := 0, len(height)-1
    max_area := 0
    for head < tail {
        area := (tail-head)*min(height[head], height[tail])
        if area > max_area {
            max_area = area
        }
        if height[head] < height[tail] {
            head++
        } else{
            tail--
        }
    }
    return max_area
}

func min(v1 int, v2 int) int {
    if v1 < v2 {
        return v1
    } else {
        return v2
    }
}
