/*
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，
另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，
后n 个元素为 0 ，应忽略。nums2 的长度为 n

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

//逻辑错误
//func merge(nums1 []int, m int, nums2 []int, n int) []int {
//	if m+n == n {
//		nums1 = nums2
//		return nums1
//	} else if n == 0 {
//		return nums1
//	}
//	newNums := make([]int, 0)
//	i := 0
//	for j := 0; j < m+n; j++ {
//		if nums1[j] < nums2[i] {
//			newNums = append(newNums, nums1[j])
//
//		} else if nums1[j] > nums2[i] {
//			newNums = append(newNums, nums2[i])
//			i += 1
//		} else {
//			newNums = append(newNums, nums1[j])
//			newNums = append(newNums, nums2[i])
//			i += 1
//		}
//		if i == n && j < m {
//			for j < m {
//				newNums = append(newNums, nums1[j])
//			}
//			break
//		}
//		if i == n && j == m {
//			break
//		}
//		if j == m-1 && i < n {
//			for i < n {
//				newNums = append(newNums, nums2[i])
//				i += 1
//			}
//			break
//		}
//	}
//	nums1 = newNums
//	return nums1
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	if m == 0 {
		return
	}

	//选择排序
	for i := 0; i < m+n-1; i++ {
		minNumIndex := i
		for j := i + 1; j < m+n; j++ {
			if nums1[j] < nums1[minNumIndex] {
				minNumIndex = j
			}
		}
		nums1[i], nums1[minNumIndex] = nums1[minNumIndex], nums1[i]
	}

}
func main() {
	//nums1 := []int{1, 2, 8, 0, 0, 0}
	//nums2 := []int{2, 5, 6}

	//fmt.Println(merge(nums1, 3, nums2, 3))
	//fmt.Println(nums1)
}
