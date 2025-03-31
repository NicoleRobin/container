package main

import (
	"fmt"
	"github.com/nicolerobin/container/union_find"
)

func main() {
	unionFind := union_find.NewQuickUnion(10)
	unionFind.Union(1, 2)
	unionFind.Union(1, 3)
	fmt.Println(unionFind)
	fmt.Println(unionFind.Find(1))
}
