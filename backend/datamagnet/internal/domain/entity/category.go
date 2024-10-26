package entity

const (
	CategoryKisi Category = "KISI"
	CategoryPdk  Category = "PDK"
)

type Category string

func (c Category) String() string {
	return string(c)
}

func (c Category) IsPreserved() bool {
	switch c {
	case CategoryKisi, CategoryPdk:
		return true
	default:
		return false
	}
}
