package anonymous

func Intseq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}