package news

type Domain struct {
	Article interface{}
}

type Repository interface {
	GetNewsByCategory(category string) (Domain, error)
}
