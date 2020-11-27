package main

import (
	"sort"
	"strings"
)

// 数组

//350.两个数组的交集 II
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
func intersect(nums1 []int, nums2 []int) []int {

	res := make([]int, 0)
	map1 := make(map[int]int, 0)
	for _,v := range nums1 {
		map1[v] += 1
	}

	for _, v := range nums2 {
		if value ,ok := map1[v]; ok&&value>0 {
			res = append(res, v)
			map1[v] -= 1
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

	base := strs[0]
	for _, v := range strs {
		for strings.Index(v, base) != 0 {
			if base == "" {
				return ""
			}
			base = base[:len(base)-1]
		}
	}
	return base
}

// 122. 买卖股票的最佳时机 II
//输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
func maxProfit(prices []int) int {
	res := 0

	for k,v := range prices {
		if k+1 < len(prices) && v < prices[k+1] {
			res += prices[k+1] - v
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
	rotateA(nums)
	rotateA(nums[:k%len(nums)])
	rotateA(nums[k%len(nums):])
}

func rotateA(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

//27. 移除元素
//给定 nums = [3,2,2,3], val = 3,
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
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
			digits[i]++
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
		temp := nums[i]
		for j := i+1; j < len(nums); j++ {
			if temp + nums[j] == target {
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
	sort.Ints(nums)

	res := make([][]int, 0)

	for first := 0; first < len(nums); first++ {

		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		target := nums[first] * -1
		third := len(nums)-1

		for second := first+1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second] + nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second] + nums[third] == target {
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

	if numRows == 1 {
		return s
	}

	d := []rune(s)
	res := make([]string, 0)
	permod := numRows*2-2

	for i := 0; i < len(d); i++ {
		mod := i % permod
		if mod < numRows {
			res[mod] += string(d[i])
		} else {
			res[permod-mod] += string(d[i])
		}
	}
	return strings.Join(res, "")
}

func main() {
	intersect([]int{1,2,2,1},[]int{2,2})
	//longestCommonPrefix([]string{"flower","flow","flight"})
	//maxProfit([]int{7,1,5,3,6,4})
	//rotate([]int{1,2,3,4,5,6,7}, 3)
	//removeElement([]int{3,2,2,3}, 3)
	//plusOne([]int{1,2,3})
	//twoSum([]int{2, 7, 11, 15}, 9)
	//threeSum([]int{-1, 0, 1, 2, -1, -4})
	//convert("LEETCODEISHIRING", 4)
}
