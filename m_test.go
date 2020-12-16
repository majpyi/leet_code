package leet_code

import (
	"fmt"
	"testing"
)

func TestM(t *testing.T) {
	fmt.Println(len("mjy"))
	for _, i := range "mjy" {
		fmt.Println(i)
	}
}

func TestGenerate(t *testing.T) {
	//fmt.Println(generate(5))
	//fmt.Println(generate(1))
	//fmt.Println(generate(2))
	//arr := make([]int, 0, 10)
	//fmt.Println(len(arr))
	//fmt.Println(cap(arr))
	fmt.Println(splitIntoFibonacci("124"))
}

func TestFib(t *testing.T) {
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
}
