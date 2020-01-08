package groph

var (
	_ RGbool      = Tree{}
	_ OutLister   = Tree{}
	_ EdgeLister  = Tree{}
	_ RootsLister = Tree{}
)
