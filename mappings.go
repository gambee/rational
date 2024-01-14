package rational

func Map(f func(Number) Number, nums []Number) (res []Number) {
    res = make([]Number, len(nums))
    for i, n := range nums {
        res[i] = f(n)
    }
    return
}
