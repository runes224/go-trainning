package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	sum1 := 0
	sum2 := 0
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	sum1 += <-ch1
	sum2 += <-ch2
	return sum1 == sum2
}

func main() {
	// ch := make(chan int)
	//go Walk(tree.New(1), ch)
	//for i := 0; i < 10; i++ {
	//	n := <-ch
	//	fmt.Print(n, ",")
	//}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}