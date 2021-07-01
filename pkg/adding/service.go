package adding

type Service interface {
	AddRestaurant(r *Restaurant) (*int, error)
}

type service struct {
	Repository Repository
}

func (s service) AddRestaurant(r *Restaurant) (*int, error) {
	return s.Repository.CreateRestaurant(r)
}

func New(repo Repository) Service {
	return service{
		Repository: repo,
	}
}
