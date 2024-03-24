package util

import (
	"fmt"
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	keys := GenerateKeys("blue", 0)
	if len(keys) != 16 {
		t.Fatal("结果错误")
	}
	fmt.Println(keys)

}
func TestGenerateKeys2(t *testing.T) {
	keys := GenerateKeys("red", 1)
	if len(keys) != 33 {
		t.Fatal("结果错误")
	}
	fmt.Println(keys)
}
func TestGenerateKeys3(t *testing.T) {
	keys := GenerateKeys("red", 2)
	sum := 33 * 32 / 2
	if len(keys) != sum {
		t.Fatal("结果错误")
	}
	fmt.Println(keys)
}
