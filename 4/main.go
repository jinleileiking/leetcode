package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if findMedian(nums1) == 0 {
		return findMedian(nums2)
	}
	if findMedian(nums2) == 0 {
		return findMedian(nums1)
	}
	return (findMedian(nums1) + findMedian(nums2)) / 2
}

func findMedian(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	// firstFloat := []float64{}
	// for i := 0; i < len(arr); i++ {
	// 	med := ((float64)(arr[i] + arr[len(arr)-1-i])) / 2
	// 	firstFloat = append(firstFloat, med)
	// }

	// floating

	firstFloat := []float64{}
	for _, v := range arr {
		firstFloat = append(firstFloat, (float64)(v))
	}

	tmpArr := firstFloat
	for {
		// spew.Dump(tmpArr)
		tmpArr = findMedianFloat(tmpArr)
		if len(tmpArr) == 1 {
			return tmpArr[0]
		}
	}
}

func findMedianFloat(arr []float64) []float64 {
	tmpArr := []float64{}
	for i := 0; i < len(arr); i++ {
		// spew.Dump(i, tmpArr)
		if i == len(arr)/2 {
			if len(arr)%2 != 0 {
				tmpArr = append(tmpArr, arr[i])
			}

			break
		}
		med := ((float64)(arr[i] + arr[len(arr)-1-i])) / 2
		tmpArr = append(tmpArr, med)
	}

	return tmpArr
}
