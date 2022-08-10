package neetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	// fixed size sliding window of size len(s1)
	// count characters of s1
	// count characters of s2 up to len(s1)
	// while r < len(s2)
	// check if it is a permutation of s1
	// return if true
	// change map by decrementing left char count, and add right char counter
	// increment both left and right index
	// check again if is permutation of s1
	// return false
	c1 := make(map[int32]int)
	for _, char := range s1 {
		c1[char] += 1
	}
	c2 := make(map[int32]int)
	for i := 0; i < len(s1); i++ {
		c2[[]rune(s2)[i]] += 1
	}

	l, r := 0, len(s1)
	for r < len(s2) {
		if isPermutation(c1, c2) {
			return true
		}
		c2[int32(s2[l])] -= 1
		if c2[int32(s2[l])] == 0 {
			delete(c2, int32(s2[l]))
		}
		c2[int32(s2[r])] += 1
		r += 1
		l += 1
	}
	if isPermutation(c1, c2) {
		return true
	}

	return false
}

func isPermutation(c1 map[int32]int, c2 map[int32]int) bool {
	if len(c1) != len(c2) {
		return false
	}
	for k, v := range c1 {
		if c2[k] != v {
			return false
		}
	}

	return true
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	values := new([]int)
	inOrder(root, values)
	return (*values)[k-1]

}

func inOrder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inOrder(root.Right, vals)
}
