package repos

type Repositories struct {
	Example IExampleRepo
}

type IExampleRepo interface {
	Create(name string) string
	Update()
	Delete()
	GetList()
}
