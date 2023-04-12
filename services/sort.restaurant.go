package services

import (
	"restaurant/models"
)

func GetRestaurantRecommendations(user *models.User, availableRestaurants []*models.Restaurant) []string {

	// 1 Filter
	result := (&FeaturedPrimaryCuisinePrimaryCostStrategy{}).Filter(user, availableRestaurants)

	if result == nil || len(result) < 1 {
		result = (&FeaturedPrimaryCuisineSecondryCostStrategy{}).Filter(user, availableRestaurants)
		result = append(result, (&FeaturedPrimaryCuisineSecondryCostStrategy{}).Filter(user, availableRestaurants)...)
	}

	// 2 Filter
	temp2 := (&PrimaryCuisinePrimaryCostStrategy{}).Filter(user, availableRestaurants)

	// 3 Filter
	temp3 := (&PrimaryCuisineSecondryCostStrategy{}).Filter(user, availableRestaurants)

	// 4 Filter
	temp4 := (&SecondaryCuisinePrimaryCostStrategy{}).Filter(user, availableRestaurants)

	result = uniqueMerge(result, temp2, temp3, temp4)
	// 5 Filter
	temp5 := (&NewltRestaurantStrategy{}).Filter(user, availableRestaurants)
	currentLen := len(result)
	result = uniqueMerge(result, temp5)
	fixLen := len(result) - currentLen
	if fixLen > 4 {
		fixLen = 4
	}
	result = result[:currentLen+fixLen]
	// 6 Filter
	temp6 := (&LowPrimaryCuisinePrimaryCostStrategy{}).Filter(user, availableRestaurants)

	// 7 Filter
	temp7 := (&LowPrimaryCuisineSecondryCostStrategy{}).Filter(user, availableRestaurants)

	// 8 Filter
	temp8 := (&LowSecondaryCuisinePrimaryCostStrategy{}).Filter(user, availableRestaurants)
	result = uniqueMerge(result, temp6, temp7, temp8, availableRestaurants)
	var ids []string
	for _, r := range result {
		ids = append(ids, r.GetRestaurantID())
	}
	return ids
}

func uniqueMerge(args ...[]*models.Restaurant) []*models.Restaurant {
	exists := make(map[string]struct{})
	var temp []*models.Restaurant
	for i := 0; i < len(args); i++ {
		if args[i] == nil || len(args) == 0 {
			continue
		}

		for _, t := range args[i] {
			if _, ok := exists[t.GetRestaurantID()]; !ok {
				exists[t.GetRestaurantID()] = struct{}{}
				temp = append(temp, t)
			}
		}
	}
	return temp
}
