package service

import (
	"fmt"
	"testing"
)

func TestGenerateRangeKey(t *testing.T) {
	keys := GenerateRangeKey(9)
	for _, v := range keys {
		fmt.Println(v.Key, v.Values)
	}
}
