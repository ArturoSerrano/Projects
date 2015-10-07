package calculators

func ArraySum(val []int) int{
	l := len(val)
	resp := 0
	for i:= 0; i < l; i++ {
		resp += val[i]
	}
	
	return resp
}