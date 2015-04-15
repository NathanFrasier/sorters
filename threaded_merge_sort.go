package sorters

import "sync"

func Tmrgsort(data []int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go tmrgsortHelper(data,&wg)
	wg.Wait()
}

func tmrgsortHelper(data []int, wg *sync.WaitGroup) {
	defer wg.Done()	//call done when the method is done regardless of where the return happens
    if len(data) < 2 {
        return 
    }
    mid := len(data) / 2
    var newWG sync.WaitGroup
    newWG.Add(2)
    go tmrgsortHelper(data[:mid],&newWG)
    go tmrgsortHelper(data[mid:],&newWG)
    newWG.Wait()
    if data[mid-1] <= data[mid] {
        return 
    }
    temp := make([]int, len(data)/2+1)
    copy(temp, data[:mid])
    l, r := 0, mid
    for i := 0; ; i++ {
        if temp[l] <= data[r] {
            data[i] = temp[l]
            l++
            if l == mid {
                break
            }
        } else {
            data[i] = data[r]
            r++
            if r == len(data) {
                copy(data[i+1:], temp[l:mid])
                break
            }
        }
    }
    return
}