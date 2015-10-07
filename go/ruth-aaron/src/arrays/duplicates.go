package arrays

func RemoveDuplicates(arr []int) []int {	
	var resp []int
	l := len(arr)
	for i:=0; i < l; i++ {
		val := arr[i]
		if !Contains(resp, val){
			resp = append(resp, val)
		}
	}
	return resp
}

func Contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}