package listing

type Service interface {
	FindAll() ([]Restaurant, error)
}

type service struct {
	Repository Repository
}

func (s service) FindAll() ([]Restaurant, error) {
	return s.Repository.FindAll()
}

func New(repo Repository) Service {
	return service{
		Repository: repo,
	}
}
