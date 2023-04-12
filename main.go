package main

import (
	"fmt"
	"restaurant/models"
	"restaurant/services"
	"time"
)

func main() {
	fmt.Println("Welcome to new zen restaurant")
	test()
}

func test() {
	restaurant1 := models.NewRestaurant("1", models.SouthIndian, 1, true, 4.5, time.Now())
	restaurant2 := models.NewRestaurant("2", models.Chinese, 2, false, 4, time.Now())
	restaurant3 := models.NewRestaurant("3", models.NorthIndian, 3, false, 3.8, time.Now())

	restaurants := []*models.Restaurant{restaurant1, restaurant2, restaurant3}

	// Create sample user
	user := models.NewUser([]*models.CuisineTracking{models.NewCuisineTracking(models.SouthIndian, 1)},
		[]*models.CostTracking{models.NewCostTracking(1, 1)})

	temp := services.GetRestaurantRecommendations(user, restaurants)
	fmt.Println(temp)
}
