package section6

func Average(vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum / len(vals)
}
