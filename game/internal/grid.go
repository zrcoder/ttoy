package internal

type GridUI struct {
	Cells [][]Cell
}

type Cell struct {
	Images       []string
	BorderTop    bool
	BorderBottom bool
	BorderLeft   bool
	BorderRight  bool
}
