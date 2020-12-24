package main

// 数组

//350.两个数组的交集 II
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
func intersect(nums1 []int, nums2 []int) []int {
	return nil
}

//14. 最长公共前缀
//输入: ["flower","flow","flight"]
//输出: "fl"
func longestCommonPrefix(strs []string) string {
	return ""
}

// 122. 买卖股票的最佳时机 II
//输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
func maxProfit(prices []int) int {
	return 0
}

//189. 旋转数组
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
func rotate(nums []int, k int)  {
	return
}

//27. 移除元素
//给定 nums = [3,2,2,3], val = 3,
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	return 0
}

//66. 加一
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
func plusOne(digits []int) []int {
	return nil
}

//1. 两数之和
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
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
	return nil
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
	return 0
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
	if head == nil {

	}

	return nil
}

//21. 合并两个有序链表
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

//141. 环形链表
//给定一个链表，判断链表中是否有环。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
//如果链表中存在环，则返回 true 。 否则，返回 false 。
func hasCycle(head *ListNode) bool {
	return false
}

//2. 两数相加
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

// 动态规划

//70. 爬楼梯
//输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
func climbStairs(n int) int {
	return 0
}

//53. 最大子序和
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
func maxSubArray(nums []int) int {
	return 0
}

//300. 最长上升子序列
//输入: [10,9,2,5,3,7,101,18]
//输出: 4
//解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
func lengthOfLIS(nums []int) int {
	return 0
}


func main() {

}
