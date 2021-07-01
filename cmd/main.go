package main

import (
	"fmt"
	"log"

	"github.com/zerodois/ddd-hexa-example/pkg/adding"
	"github.com/zerodois/ddd-hexa-example/pkg/listing"
	"github.com/zerodois/ddd-hexa-example/pkg/postgres"
)

func main() {
	repo := postgres.New()
	add := adding.New(repo)
	list := listing.New(repo)

	if _, err := add.AddRestaurant(&adding.Restaurant{Name: "Restaurante Goomer"}); err != nil {
		log.Fatal(err)
	}

	if _, err := add.AddRestaurant(&adding.Restaurant{Name: "The Lícia"}); err != nil {
		log.Fatal(err)
	}

	restaurants, err := list.FindAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(restaurants)
}
