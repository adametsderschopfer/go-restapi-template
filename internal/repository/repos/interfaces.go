package repos

type Repositories struct {
	Example IExampleRepo
}

type IExampleRepo interface {
	Create(name string)
	Update()
	Delete()
	GetList()
}
