package main

import (
	"context"
	"fmt"

	"github.com/flvsantos15/imersao20/simulator/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoStr := "mongodb://admin:admin@localhost:27017/routes?authSource=admin"
	mongoConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoStr))
	if err != nil {
		panic(err)
	}

	freightService := internal.NewFreightService()
	routeService := internal.NewRouteService(mongoConnection, freightService)

	routeCreatedEvent := internal.NewRouteCreatedEvent(
		"1",
		100,
		[]internal.Directions{
			{Lat: 0, Lng: 0},
			{Lat: 10, Lng: 10},
		},
	)

	fmt.Println(internal.RouteCreatedHandler(routeCreatedEvent, routeService))
}
