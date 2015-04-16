package sorters

import "sync"

func tpartition(data []int) int {
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
 
func qcksorthelper(data []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if(len(data) <= 1) {
		return
	}
    pivotIndex := tpartition(data)
    var newWG sync.WaitGroup
    if(pivotIndex > 0) {
    	newWG.Add(1)
    	qcksorthelper(data[:pivotIndex],&newWG)
    }
    if(pivotIndex < len(data)-1) {
    	newWG.Add(1)
    	qcksorthelper(data[pivotIndex+1:],&newWG)
    }
    newWG.Wait()
}
func Tqcksort(data[]int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go qcksorthelper(data,&wg)
	wg.Wait()
}