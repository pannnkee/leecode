package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// 数组

//350.两个数组的交集 II
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
func intersect(nums1 []int, nums2 []int) []int {

	var res []int
	intMap := make(map[int]int,0)
	for _, v := range nums1 {
		intMap[v] += 1
	}

	for _, v := range nums2 {
		if i, ok := intMap[v]; ok && i > 0 {
			intMap[v] -= 1
			res = append(res, v)
		}
	}

	return res
}

//14. 最长公共前缀
//输入: ["flower","flow","flight"]
//输出: "fl"
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0{
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// 122. 买卖股票的最佳时机 II
//输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
func maxProfit(prices []int) int {

	var res int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}

//189. 旋转数组
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
func rotate(nums []int, k int)  {
	rotateArray(nums)
	rotateArray(nums[:k%len(nums)])
	rotateArray(nums[k%len(nums):])
}

func rotateArray(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

//27. 移除元素
//给定 nums = [3,2,2,3], val = 3,
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums);{
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

//66. 加一
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
func plusOne(digits []int) []int {
	for i := len(digits)-1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

//1. 两数之和
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}

//15. 三数之和
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums)-1
		target := nums[first]*-1

		for second := first+1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second] + nums[third] > target{
				third --
			}
			if second == third {
				break
			}
			if target == nums[second] + nums[third] {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

//6. Z 字形变换
//输入: s = "LEETCODEISHIRING", numRows =4
//输出:"LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
func convert(s string, numRows int) string {
	return ""
}

//268. 丢失的数字
//输入：nums = [3,0,1]
//输出：2
//解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
func missingNumber(nums []int) int {
	intMsp := make(map[int]int, 0)
	for k, v := range nums {
		intMsp[v] = k
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := intMsp[i]; !ok {
			return i
		}
	}
	return -1
}


// 链表

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

//19. 删除链表的倒数第N个节点
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newListNode := &ListNode{Next: head}
	res := newListNode

	fast, slow := newListNode, newListNode
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return res.Next
}

//21. 合并两个有序链表
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newListNode := &ListNode{}
	res := newListNode

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			newListNode.Next = l2
			l2 = l2.Next
		} else {
			newListNode.Next = l1
			l1 = l1.Next
		}
		newListNode = newListNode.Next
	}
	if l1 != nil {
		newListNode.Next = l1
	}
	if l2 != nil {
		newListNode.Next = l2
	}
	return res.Next
}

//141. 环形链表
//给定一个链表，判断链表中是否有环。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//如果链表中存在环，则返回 true 。 否则，返回 false 。
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast,slow := head.Next, head
	for fast != nil && slow != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

//2. 两数相加
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := 0
	res := head

	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{Val:  tmp%10}
		head = head.Next
		tmp = tmp / 10
	}
	return res.Next
}

// 动态规划

//70. 爬楼梯
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n-1]
}

//53. 最大子序和
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = nums[i] + dp[i-1]
		}
	}
	sort.Ints(dp)
	return dp[len(nums)-1]
}

//300. 最长上升子序列
//输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
func lengthOfLIS(nums []int) int {
	return 0
}


// 字符串
//344. 反转字符串
//输入：["h","e","l","l","o"]
//输出：["o","l","l","e","h"]
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i],s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return
}

//387. 字符串中的第一个唯一字符
//s = "leetcode"
//返回 0
//
//s = "loveleetcode"
//返回 2

func firstUniqChar(s string) int {
	var arr [26]int

	for k,v := range s {
		arr[v - 'a'] = k
	}
	for k,v := range s {
		if arr[v - 'a'] == k {
			return k
		} else {
			arr[v - 'a'] = -1
		}
	}
	return -1
}

//28. 实现 strStr()
//输入: haystack = "hello", needle = "ll"
//输出: 2
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	length := len(haystack)
	length2 := len(needle)
	for k := range haystack {
		if length2 + k < length {
			if haystack[k:k+length2] == needle {
				return k
			}
		} else {
			return -1
		}
	}
	return -1
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	fmt.Println(n)
	return fib(n-1)+fib(n-2)
}

//输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
//输入: n = 1
//输出: [1,2,3,4,5,6,7,8,9]
//剑指offer17
func printNumbers(n int) []int {
	res := make([]int,0)

	tmp := 0
	for n > 0 {
		n--
		tmp = tmp*10 + 9
	}
	for i := 1; i < tmp+1; i++ {
		res = append(res, i)
	}
	return res
}

//125. 验证回文串
//输入: "A man, a plan, a canal: Panama"
//输出: true
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	temp := []byte{}
	for _,v := range s {
		if isanum(byte(v)) {
			temp = append(temp, byte(v))
		}
	}
	for i := 0; i < len(temp)/2; i++ {
		if temp[i] != temp[len(temp)-i-1] {
			return false
		}
	}
	return true
}

func isanum(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}

//796. 旋转字符串
//给定两个字符串, A 和 B。A 的旋转操作就是将 A 最左边的字符移动到最右边。
//例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。
//输入: A = 'abcde', B = 'cdeab'
//输出: true
func rotateString(A string, B string) bool {
	if A == "" && B == "" {
		return true
	}
	for i := 0; i < len(A); i++ {
		if A[i:] + A[:i] == B {
			return true
		}
	}
	return false
}

//58. 最后一个单词的长度
//输入: "Hello World"
//输出: 5
func lengthOfLastWord(s string) int {
	res := []byte{}
	for i := len(s)-1; i >= 0 ; i-- {
		if s[i] != ' ' {
			res = append(res, s[i])
		} else {
			if len(res) == 0 {
				continue
			} else {
				break
			}
		}
	}
	return len(res)
}

//二叉树
//104. 二叉树的最大深度
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回它的最大深度 3
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right))+1
}

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//102. 二叉树的层序遍历
//    3
//   / \
//  9  20
//    /  \
//   15   7
//[
//  [3],
//  [9,20],
//  [15,7]
//]
func levelOrder(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int)
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}

//98. 验证二叉搜索树
//输入:
//    2
//   / \
//  1   3
//输出: true
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return test1(root, math.MaxInt64, math.MinInt64)
}

func test1(root *TreeNode, max, min int) bool {
	if root == nil {
		return false
	}
	if root.Val < min || root.Val > max {
		return false
	}
	return test1(root.Left, min, root.Val) && test1(root.Right, max, root.Val)
}

//700. 二叉搜索树中的搜索
//给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。
//返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
//给定二叉搜索树:
//
//        4
//       / \
//      2   7
//     / \
//    1   3
//
//和值: 2
//你应该返回如下子树:
//
//      2
//     / \
//    1   3

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}

//450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}

	minNode := root.Right
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	return root
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left,mid,right := 0, len(nums), (len(nums)-1)/2

	for right >= 1 {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (1 + right)/2
	}
	return -1
}

func main () {
	test()
}

func test() {

}







