package takeuforward

func heapFunc(arr [][]int) [][]int {
	ans := [][]int{}
	h := &heap{}
	for _, nums := range arr {
		h.add(nums)
	}
	for len(h.nums) != 0 {
		ans = append(ans, h.extract())
	}
	return ans
}

type heap struct {
	nums [][]int
}

func getParent(i int) int {
	return (i - 1) / 2
}
func getLeftChild(i int) int {
	return 2*i + 1
}
func getRightChild(i int) int {
	return 2*i + 2
}
func (h *heap) heapifyUp(i int) {
	for h.nums[i][0] < h.nums[getParent(i)][0] {
		h.nums[i], h.nums[getParent(i)] = h.nums[getParent(i)], h.nums[i]
		i = getParent(i)
	}
}
func (h *heap) heapifyDown(i int) {
	if len(h.nums) > 0 {
		lChild, rChild := getLeftChild(i), getRightChild(i)
		childToCompare := -1
		for lChild <= len(h.nums)-1 {
			if lChild == len(h.nums)-1 {
				childToCompare = lChild
			} else if h.nums[lChild][0] < h.nums[rChild][0] {
				childToCompare = lChild
			} else {
				childToCompare = rChild
			}
			if h.nums[i][0] > h.nums[childToCompare][0] {
				h.nums[i], h.nums[childToCompare] = h.nums[childToCompare], h.nums[i]
				i = childToCompare
				lChild = getLeftChild(i)
				rChild = getRightChild(i)
			} else {
				return
			}
		}
	}

}
func (h *heap) add(arr []int) {
	h.nums = append(h.nums, arr)
	h.heapifyUp(len(h.nums) - 1)
}
func (h *heap) extract() []int {
	if len(h.nums) != 0 {
		top := h.nums[0]
		if len(h.nums) == 1 {
			h.nums = [][]int{}
		} else {
			h.nums[0] = h.nums[len(h.nums)-1]
			h.nums = h.nums[:len(h.nums)-1]
			h.heapifyDown(0)
		}
		return top
	}
	return []int{}
}
