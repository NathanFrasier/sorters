package sorters

func Inssort(data []int) {
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