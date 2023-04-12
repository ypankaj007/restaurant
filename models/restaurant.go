package models

import (
	"time"
)

type Cuisine string

const (
	SouthIndian Cuisine = "SouthIndian"
	NorthIndian Cuisine = "NorthIndian"
	Chinese     Cuisine = "Chinese"
)

type Restaurant struct {
	restaurantID  string
	cuisine       Cuisine
	costBracket   int
	isRecommended bool
	rating        float32
	onboardedTime time.Time
}

func NewRestaurant(restaurantID string, cuisine Cuisine, costBracket int, isRecommended bool, rating float32, onboardedTime time.Time) *Restaurant {
	return &Restaurant{restaurantID: restaurantID, cuisine: cuisine, costBracket: costBracket, isRecommended: isRecommended, rating: rating, onboardedTime: onboardedTime}
}

func (restaurant *Restaurant) GetRestaurantID() string {
	return restaurant.restaurantID
}

func (restaurant *Restaurant) GetCuisine() Cuisine {
	return restaurant.cuisine
}

func (restaurant *Restaurant) GetCostBracket() int {
	return restaurant.costBracket
}

func (restaurant *Restaurant) GetIsRecommended() bool {
	return restaurant.isRecommended
}

func (restaurant *Restaurant) GetRating() float32 {
	return restaurant.rating
}

func (restaurant *Restaurant) GetOnboardedTime() time.Time {
	return restaurant.onboardedTime
}

func (restaurant *Restaurant) SetRestaurantID(restaurantID string) *Restaurant {
	restaurant.restaurantID = restaurantID
	return restaurant
}

func (restaurant *Restaurant) SetCuisine(cuisine Cuisine) *Restaurant {
	restaurant.cuisine = cuisine
	return restaurant
}

func (restaurant *Restaurant) SetCostBracket(costBracket int) *Restaurant {
	restaurant.costBracket = costBracket
	return restaurant
}

func (restaurant *Restaurant) SetIsRecommended(isRecommended bool) *Restaurant {
	restaurant.isRecommended = isRecommended
	return restaurant
}

func (restaurant *Restaurant) SetRating(rating float32) *Restaurant {
	restaurant.rating = rating
	return restaurant
}

func (restaurant *Restaurant) SetOnboardedTime(onboardedTime time.Time) *Restaurant {
	restaurant.onboardedTime = onboardedTime
	return restaurant
}

// sorting
type NewlyRest struct {
	Restaurants []*Restaurant
}

func (n *NewlyRest) Len() int {
	return len(n.Restaurants)
}
func (n *NewlyRest) Less(i, j int) bool {

	return n.Restaurants[i].GetOnboardedTime().Before(n.Restaurants[j].GetOnboardedTime())
}

func (n *NewlyRest) Swap(i, j int) {
	n.Restaurants[i], n.Restaurants[j] = n.Restaurants[j], n.Restaurants[i]
}
