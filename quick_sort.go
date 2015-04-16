package sorters

func partition(data []int,) int {
	pivotIndex := len(data)-1
	if(pivotIndex == 1) { //if length is 2 
		if(data[0] > data[pivotIndex]) {
			data[0], data[pivotIndex] = data[pivotIndex], data[0]
			pivotIndex = 0
		}
		return pivotIndex
	}
    left := 0
    right := len(data)-2
    for left < right{
        for left < len(data)-1 && data[left] <= data[pivotIndex]{
            left++
        }
        for right > 0 && data[right] >= data[pivotIndex] {
            right--
        }
        if left < right {
            data[left], data[right] = data[right], data[left]
            left++
            right--
        }
    }
    if left == right && data[left] < data[pivotIndex] { //edge case to do with odd swap counts
    	left++
    }
    data[left], data[pivotIndex] = data[pivotIndex], data[left]
    return left
}
 
func Qcksort(data []int) {
	if(len(data) <= 1) {
		return
	}
    pivotIndex := partition(data)
    if(pivotIndex > 0) {
    	Qcksort(data[:pivotIndex])
    }
    if(pivotIndex < len(data)-1) {
    	Qcksort(data[pivotIndex+1:])
    }
}