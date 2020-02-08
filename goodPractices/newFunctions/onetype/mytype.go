package onetype

type OneType struct {
	Size int
}

func New() OneType {
	return OneType{Size: 100}
}
