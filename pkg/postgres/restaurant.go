package postgres

import (
	"encoding/json"

	"github.com/zerodois/ddd-hexa-example/pkg/adding"
	"github.com/zerodois/ddd-hexa-example/pkg/listing"
)

type Postgres interface {
	CreateRestaurant(r *adding.Restaurant) (*int, error)
	FindAll() ([]listing.Restaurant, error)
}

type inMemory struct {
	Restaurants []string
}

func (m *inMemory) CreateRestaurant(r *adding.Restaurant) (*int, error) {
	serial, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	m.Restaurants = append(m.Restaurants, string(serial))
	s := len(m.Restaurants)
	return &s, nil
}

func (m *inMemory) FindAll() ([]listing.Restaurant, error) {
	rr := []listing.Restaurant{}
	for _, r := range m.Restaurants {
		restaurant := listing.Restaurant{}
		if err := json.Unmarshal([]byte(r), &restaurant); err != nil {
			return nil, err
		}
		rr = append(rr, restaurant)
	}
	return rr, nil
}

func New() Postgres {
	return &inMemory{}
}
