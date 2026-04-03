type entry struct {
    num int
    count int
}

func topKFrequent(nums []int, k int) []int {
    data := make(map[int]int)
    numsLen := len(nums)
    for i := 0; i < numsLen; i++ {
        data[nums[i]]++
    }

    entries := make([]entry, 0)
    for k, v := range data {
        entries = append(entries, entry{
            num: k,
            count: v,
        })
    }

    entries = entriesSort(entries)
    entriesLen := len(entries)
    result := make([]int, 0)
    for i := 0; i < entriesLen; i++ {
        result = append(result, entries[i].num)
    }

    return result[len(result)-k:]
}

func entriesSort(entries []entry) []entry {
    entriesLen := len(entries)
    
    if entriesLen < 2 {
        return entries
    }

    piv := entries[entriesLen / 2].count
    left := []entry{}
    mid := []entry{}
    right := []entry{}

    for i := 0; i < entriesLen; i++ {
        if piv > entries[i].count {
            left = append(left, entries[i])    
        } else if piv == entries[i].count {
            mid = append(mid, entries[i])
        } else {
            right = append(right, entries[i])
        }
    }

    res := append(entriesSort(left), mid...)
    return append(res, entriesSort(right)...)
}




