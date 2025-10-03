package sort

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	} else {
		mid := n / 2
		left := MergeSort(arr[:mid])
		right := MergeSort(arr[mid:])
		return merge(left, right)
	}
}

func merge(a, b []int) []int {
	var sorted []int
	i, j := 0, 0; 
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}

	sorted = append(sorted, a[i:]...)
	sorted = append(sorted, b[j:]...)
	return sorted
}