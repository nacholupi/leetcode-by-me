package array

import "sort"

func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	r := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {

		if nums1[i] == nums2[j] {
			r = append(r, nums1[i])
			i++
			j++
			continue
		}

		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return r
}

func Intersect2(nums1 []int, nums2 []int) []int {
	r := []int{}
	for _, v := range nums1 {
		for _, v2 := range nums2 {
			if v == v2 {
				r = append(r, v)
				break
			}
		}
	}

	return r
}
