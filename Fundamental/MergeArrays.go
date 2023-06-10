package Fundamental

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func MergeArrays(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	if arr1[0] > arr1[len(arr1)-1] {
		reverse(arr1)
	}

	if arr2[0] > arr2[len(arr2)-1] {
		reverse(arr2)
	}

	i, j := 0, 0
	var res []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			if len(res) == 0 || arr1[i] > res[len(res)-1] {
				res = append(res, arr1[i])
			}
			i++
		} else if arr1[i] == arr2[j] {
			if len(res) == 0 || arr1[i] > res[len(res)-1] {
				res = append(res, arr1[i])
			}
			i++
			j++
		} else {
			if len(res) == 0 || arr2[j] > res[len(res)-1] {
				res = append(res, arr2[j])
			}
			j++
		}
	}
	for i < len(arr1) {
		if arr1[i] > res[len(res)-1] {
			res = append(res, arr1[i])
		}
		i++

	}
	for j < len(arr2) {
		if arr2[j] > res[len(res)-1] {
			res = append(res, arr2[j])
		}
		j++
	}
	return res
}

/*func MergeArrays(arr1, arr2 []int) []int {
	answer := []int{}
	combined := append(arr1, arr2...)
	sort.Ints(combined)
	for index, item := range combined {
		if index == 0 {
			answer = append(answer, item)
			continue
		}
		if item > combined[index-1] {
			answer = append(answer, item)
		}
	}
	return answer
}*/
