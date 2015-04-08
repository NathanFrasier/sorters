package sorters

func Tmrgsort(data []int) {
    if len(data) < 2 {
        return 
    }
    mid := len(data) / 2
    Mrgsort(data[:mid])
    Mrgsort(data[mid:])
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