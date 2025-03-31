package union_find

import "fmt"

/*
QuickUnion is a simple implementation of union-find algorithm.
*/
type QuickUnion struct {
	size int
	fa   []int
	rank []int
}

func NewQuickUnion(size int) *QuickUnion {
	fa := make([]int, size)
	for i := range fa {
		fa[i] = i
	}

	rank := make([]int, size)
	for i := range rank {
		rank[i] = 1
	}
	return &QuickUnion{size: size, fa: fa}
}

func (qu *QuickUnion) String() string {
	return fmt.Sprintf("fa:%+v", qu.fa)
}

func (qu *QuickUnion) Union(x, y int) {
	xRoot := x
	for qu.fa[xRoot] != xRoot {
		xRoot = qu.fa[xRoot]
	}
	yRoot := y
	for qu.fa[yRoot] != yRoot {
		yRoot = qu.fa[yRoot]
	}

	qu.fa[xRoot] = yRoot
}

// UnionByRank 按深度合并
func (qu *QuickUnion) UnionByRank(x, y int) {
	xRoot := qu.Find(x)
	yRoot := qu.Find(y)
	if xRoot == yRoot {
		return
	}

	if qu.rank[xRoot] < qu.rank[yRoot] {
		qu.fa[xRoot] = yRoot
	} else if qu.rank[yRoot] > qu.rank[xRoot] {
		qu.fa[yRoot] = xRoot
	} else {
		qu.fa[yRoot] = xRoot
		qu.rank[xRoot] += 1
	}
}

// UnionBySize 按大小合并
func (qu *QuickUnion) UnionBySize(x, y int) {
	xRoot := qu.Find(x)
	yRoot := qu.Find(y)
	if xRoot == yRoot {
		return
	}
}

/*
在集合很大或者树不平衡时，快速合并最坏情况会退化成一条链，查询的时间复杂度为O(n)，这种情况下需要使用路径压缩进行优化；

路径压缩（Path Compression）：
在从底向上查找根节点过程中，如果此时访问的节点不是根节点，则我们可以把这个节点尽量向上移动一下，从而减少树的层树。这个过程就叫做路径压缩。

路径压缩有两种方式：
一种叫做「隔代压缩」: 也就是只压缩每隔一个节点，这样可以减少时间复杂度，但是压缩的效果不是很好；
另一种叫做「完全压缩」: 也就是每次都把当前节点的父节点指向根节点，这样虽然时间复杂度会增加，但是压缩的效果会更好。

因为路径压缩只在查询时进行，并且只压缩一棵树上的路径，所以并查集最终的结构仍然可能是比较复杂的。为了避免这种情况，另一个优化方式是「按秩合并」。
按秩合并（Union By Rank）：指的是在每次合并操作时，都把「秩」较小的树根节点指向「秩」较大的树根节点。
这里的「秩」有两种定义，一种定义指的是树的深度；另一种定义指的是树的大小（即集合节点个数）。无论采用哪种定义，集合的秩都记录在树的根节点上。

按秩合并也有两种方式：一种叫做「按深度合并」；另一种叫做「按大小合并」。
*/

func (qu *QuickUnion) Find(x int) int {
	xRoot := x
	for qu.fa[xRoot] != xRoot {
		xRoot = qu.fa[xRoot]
	}

	return xRoot
}

// Find1 隔代压缩
func (qu *QuickUnion) Find1(x int) int {
	for qu.fa[x] != x {
		qu.fa[x] = qu.fa[qu.fa[x]]
		x = qu.fa[x]
	}

	return x
}

// Find2 完全压缩
func (qu *QuickUnion) Find2(x int) int {
	if qu.fa[x] != x {
		qu.fa[x] = qu.Find2(qu.fa[x])
	}
	return x
}

func (qu *QuickUnion) IsConnected(x, y int) bool {
	return qu.Find(x) == qu.Find(y)
}
