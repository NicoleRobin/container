package union_find

/*
QuickFind is a simple implementation of UnionFind interface.
查询时间复杂度：O(1)
合并时间复杂度：O(n)
*/
type QuickFind struct {
	size int
	ids  []int
}

func NewQuickFind(size int) *QuickFind {
	ids := make([]int, size)
	for i := range ids {
		ids[i] = i
	}

	return &QuickFind{
		size: size,
		ids:  ids,
	}
}

func (qf *QuickFind) Find(x int) int {
	if x < 0 || x >= qf.size {
		return -1
	}
	return qf.ids[x]
}

func (qf *QuickFind) Union(x, y int) {
	if x < 0 || x >= qf.size || y < 0 || y >= qf.size {
		return
	}

	xId := qf.Find(x)
	yId := qf.Find(y)
	if xId != yId {
		for i, id := range qf.ids {
			if id == yId {
				qf.ids[i] = xId
			}
		}
	}
}
