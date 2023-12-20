package link

type Link struct {
	Root      string
	ChildLink []string
}

func NewLink(root string) *Link {
	return &Link{Root: root}
}
