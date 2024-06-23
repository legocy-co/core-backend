package helpers

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	pwd, err := HashPassword("12341234")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(pwd))
}
