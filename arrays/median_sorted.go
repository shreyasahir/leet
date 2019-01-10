package main

import (
	"fmt"
)

func main() {
	nums1 := []int{5}
	nums2 := []int{1, 2, 3, 4, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	n1 := len(nums1)
	n2 := len(nums2)
	if n1 == 0 {
		if n2%2 == 0 {
			return float64((nums2[n2/2] + nums2[n2/2-1])) / 2.0
		}
		return float64(nums2[n2/2])

	}
	if n2 == 0 {
		if n1%2 == 0 {
			return float64(nums1[n1/2]+nums1[(n1)/2-1]) / 2.0
		}
		return float64(nums1[n1/2])

	}
	var result []int
	n := n1 + n2
	even := false
	if n%2 == 0 {
		even = true
	}
	n = n / 2
	for i, j := 0, 0; i < n1 && j < n2; {
		//fmt.Println("value of n", n)
		//fmt.Println("result", result)
		if len(result) == n+1 {
			if even {
				//fmt.Println("i am here", float64((result[n] + result[n-1])))
				return float64((result[n] + result[n-1])) / 2.0
			}
			return float64(result[n])
		}
		if nums1[i] > nums2[j] {
			result = append(result, nums2[j])
			j++
		} else if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			result = append(result, nums2[j])
			i++
			j++
		}
		if i == len(nums1) {
			result = append(result, nums2[j:]...)
		}
		if j == len(nums2) {
			result = append(result, nums1[i:]...)
		}

	}
	//fmt.Println("result_final", result)

	if even {
		//fmt.Println("its even")
		return float64((result[n] + result[n-1])) / 2.0
	}
	return float64(result[n])
}
