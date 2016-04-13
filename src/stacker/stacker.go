package main

import (
	"fmt"
	"stacker/stack"
)

func main(){
	var haystack stack.Stack
	haystack.Push("Remco")
	haystack.Push(43.4)
	haystack.Push(-120)
	haystack.Push([]string{"willem","pieter"})
	haystack.Push([]int{7,5})
	haystack.Push("Clara")
	haystack.Push(40)

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}



