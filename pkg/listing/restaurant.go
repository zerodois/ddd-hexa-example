package listing

type Restaurant struct {
	Name string `json:"name"`
}

type Repository interface {
	FindAll() ([]Restaurant, error)
}
