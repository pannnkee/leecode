package main

// 数组

//350.两个数组的交集 II
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
func intersect(nums1 []int, nums2 []int) []int {

}

//14. 最长公共前缀
//输入: ["flower","flow","flight"]
//输出: "fl"
func longestCommonPrefix(strs []string) string {

}

// 122. 买卖股票的最佳时机 II
//输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
func maxProfit(prices []int) int {

}

//189. 旋转数组
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
func rotate(nums []int, k int)  {

}

//27. 移除元素
//给定 nums = [3,2,2,3], val = 3,
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {

}

//66. 加一
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
func plusOne(digits []int) []int {

}

//1. 两数之和
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
func twoSum(nums []int, target int) []int {

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

}

func main() {
	intersect([]int{1,2,2,1},[]int{2,2})
	longestCommonPrefix([]string{"flower","flow","flight"})
	maxProfit([]int{7,1,5,3,6,4})
	rotate([]int{1,2,3,4,5,6,7}, 3)
	removeElement([]int{3,2,2,3}, 3)
	plusOne([]int{1,2,3})
	twoSum([]int{2, 7, 11, 15}, 9)
	threeSum([]int{-1, 0, 1, 2, -1, -4})
	convert("LEETCODEISHIRING", 4)
}