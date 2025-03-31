package union_find

type UnionFind interface {
	Union(x, y int)
	Find(x int) int
	IsConnected(x, y int) bool
}
