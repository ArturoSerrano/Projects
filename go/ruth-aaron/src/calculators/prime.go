package calculators

import (
	"fmt"
)
func CalcPrimes(val int) []int{
	var resp []int 
	mid := val/2
	for i:=2; i < mid; i++ {
		if val%i == 0{
			q := val/i
			
			if IsPrime(q) {
				fmt.Printf("q is prime: %d\n", q)
				resp = append(resp, q)
			} else {
				resp = append(resp, CalcPrimes(q)...)
			}
			
			if IsPrime(i) {
				fmt.Printf("i is prime: %d\n", i)
				resp = append(resp, i)
			} else {
				resp = append(resp, CalcPrimes(i)...)	
			}
			
		} 	
	}
	
	return resp
}


func IsPrime(val int) bool {
	if val == 2 {
		return true
	}
	
	mid := val/2
	for i:=2; i <= mid; i++ {
		if val%i == 0{
			return false;
		}
	}
	
	return true
}