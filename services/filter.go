package services

import (
	"log"
	"restaurant/models"
	"sort"
)

type Critaria interface {
	Filter([]*models.Restaurant) []*models.Restaurant
}

type CuisineFilter struct {
	Cuisine models.Cuisine
}

func (cf *CuisineFilter) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	var temp []*models.Restaurant
	for _, restaurant := range restaurants {
		if restaurant.GetCuisine() == cf.Cuisine {
			temp = append(temp, restaurant)
		}
	}
	return temp
}

// type SecondryCuisine struct {
// 	Cuisine models.Cuisine
// }

// func (sc *SecondryCuisine) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
// 	return []*models.Restaurant{}
// }

type CostFilter struct {
	Cost int
}

func (cf *CostFilter) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	var temp []*models.Restaurant
	for _, restaurant := range restaurants {
		if restaurant.GetCostBracket() == cf.Cost {
			temp = append(temp, restaurant)
		}
	}
	return temp
}

// type SecondryCost struct {
// 	Cost float32
// }

// func (sc *SecondryCost) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
// 	return []*models.Restaurant{}
// }

type GreaterRating struct {
	Rating float32
}

func (gr *GreaterRating) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	var temp []*models.Restaurant
	for _, restaurant := range restaurants {
		if restaurant.GetRating() >= gr.Rating {
			temp = append(temp, restaurant)
		}
	}
	return temp
}

type LessRating struct {
	Rating float32
}

func (lr *LessRating) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	var temp []*models.Restaurant
	for _, restaurant := range restaurants {
		if restaurant.GetRating() < lr.Rating {
			temp = append(temp, restaurant)
		}
	}
	return temp
}

type AndFilter struct {
	FirstCritaria  Critaria
	SecondCritaria Critaria
}

func (af *AndFilter) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	temp := af.FirstCritaria.Filter(restaurants)
	log.Println(temp)
	temp = af.SecondCritaria.Filter(temp)
	log.Println(temp)
	return temp
}

type OrFilter struct {
	FirstCritaria  Critaria
	SecondCritaria Critaria
}

func (of *OrFilter) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	temp := of.FirstCritaria.Filter(restaurants)
	temp = append(temp, of.SecondCritaria.Filter(restaurants)...)
	return temp
}

type FeaturedrFilter struct {
}

func (ff *FeaturedrFilter) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	var temp []*models.Restaurant
	for _, restaurant := range restaurants {
		if restaurant.GetIsRecommended() {
			temp = append(temp, restaurant)
		}
	}
	return temp
}

type NewlyFilter struct {
}

func (nf *NewlyFilter) Filter(restaurants []*models.Restaurant) []*models.Restaurant {
	var temp []*models.Restaurant
	copy(temp, restaurants)
	var obj = &models.NewlyRest{Restaurants: temp}
	sort.Sort(obj)
	return obj.Restaurants
}
