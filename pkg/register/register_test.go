package register

import (
	"fmt"
	"testing"
)

func TestGenerateActivationCode(t *testing.T) {
	data := GenerateActivationCode("0")
	ac, ok := ValidateActivationCode(data)
	if ok {
		fmt.Printf("regist ok : %#v", ac)
		return
	}
	fmt.Println("regist failed ")
}
