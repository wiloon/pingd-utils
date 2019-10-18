package utils

import (
	"fmt"
	"testing"
)

func TestStringToDate(t *testing.T) {
	str := "Fri, 18 Oct 2019 07:40:06 GMT"
	foo := StringToDate(str)
	fmt.Println(foo)
}
