package validators
import (
	"strconv"
)
func IsNum (variable string) (value int, err bool)  {
	if value, err := strconv.Atoi(variable); err == nil {
		return value, false 
	} 
	return -1, true
}