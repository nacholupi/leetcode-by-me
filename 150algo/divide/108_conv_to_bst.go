package divide

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTRec(nums, 0, len(nums)-1)
}

func sortedArrayToBSTRec(nums []int, start int, end int) *TreeNode {

	if end-start < 0 {
		return nil
	}

	mid := start + (end-start)/2

	res := &TreeNode{Val: nums[mid]}

	res.Left = sortedArrayToBSTRec(nums, start, mid-1)

	res.Right = sortedArrayToBSTRec(nums, mid+1, end)

	return res
}
