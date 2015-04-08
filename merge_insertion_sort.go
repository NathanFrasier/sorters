package sorters

func Mrginssort(data []int) {
    if len(data) < 100 {
        //implement insertion sort here 
        for i := 1; i < len(data); i++ {
        value := data[i]
        j := i - 1
        for j >= 0 && data[j] > value {
            data[j+1] = data[j]
            j = j - 1
        }
        data[j+1] = value
    }
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