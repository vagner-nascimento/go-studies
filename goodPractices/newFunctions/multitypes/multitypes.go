package multitypes

type One struct {
	Name string
}

type Two struct {
	Name string
}

func NewOne() One {
	return One{Name: "One"}
}

func NewTwo() Two {
	return Two{Name: "Two"}
}
