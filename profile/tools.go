package profile

func Avg(items []VoteFloat) float64 {
	var tot int64 = 0
	for _, k := range items {
		tot += int64(k.Value)
	}
	return float64(tot / int64(len(items)))
}

func Avg16(items []VoteInt) float64 {
	var tot int64 = 0
	for _, k := range items {
		tot += int64(k.Value)
	}
	return float64(tot / int64(len(items)))
}
