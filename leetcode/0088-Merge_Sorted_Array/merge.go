package test

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
    if  n > 0 {
		if m == 0 {
			for i, v := range nums2 {
				nums1[i] = v
			}
		} else {
			c := n-1
			t := m+n-1
			for i := m-1; i >= 0; i-- {
				for j := c; j >= 0; j-- {
					if nums2[j] > nums1[i] {
						nums1[t] = nums2[j]
						fmt.Println(t, nums1)
						c--
						t--
					} else {
						nums1[t] = nums1[i]
						fmt.Println(t, nums1)
						t--
						break
					}
				}
			}
			if c >= 0 {
				for j := c; j >= 0; j-- {
					nums1[t] = nums2[j]
					fmt.Println(t, nums1)
					c--
					t--
				}
			}
		}
	}
}
