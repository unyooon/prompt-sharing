package types

type QueryInterface interface {
	GetPage() int
	SetPage(page int)
	GetSize() int
	SetSize(size int)
}

type Query struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (q *Query) GetPage() int {
	return q.Page
}

func (q *Query) SetPage(page int) {
	q.Page = page
}

func (q *Query) GetSize() int {
	return q.Size
}

func (q *Query) SetSize(size int) {
	q.Size = size
}
