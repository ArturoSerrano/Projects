package main

import (
	"fmt"
	"os"
	"validators"
	"calculators"
	"arrays"
)

func main() {
	fmt.Printf ("Args: %q\n", os.Args)
	val1,err1 := validators.IsNum(os.Args[1])
	val2,err2 := validators.IsNum(os.Args[2])
	
	if err1 || err2 {
		fmt.Printf("Not valid Entries")
		return
	}
	
	if val1 - val2 != 1 || val2 - val1 != -1 {
		fmt.Printf("Not Consecutive")
		return
	}
	
	primeSum1 :=calculators.ArraySum(arrays.RemoveDuplicates(calculators.CalcPrimes(val1)))
	primeSum2 :=calculators.ArraySum(arrays.RemoveDuplicates(calculators.CalcPrimes(val2)))
	
	if primeSum1 == primeSum2 {
		fmt.Printf("VALID: %d = %d", primeSum1, primeSum2)
	} else {
		fmt.Printf("INVALID: %d /= %d", primeSum1, primeSum2)
	}
}



