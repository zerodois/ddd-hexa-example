package adding

type Restaurant struct {
	Name string `json:"name"`
}

type Repository interface {
	CreateRestaurant(r *Restaurant) (*int, error)
}
