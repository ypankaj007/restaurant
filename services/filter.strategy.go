package services

import (
	"log"
	"restaurant/models"
)

type FilterStrategy interface {
	Filter(*models.User, []*models.Restaurant) []*models.Restaurant
}

type FeaturedPrimaryCuisinePrimaryCostStrategy struct{}

func (pc *FeaturedPrimaryCuisinePrimaryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimryCost() || !user.HasPrimaryCuisine() {
		var t []*models.Restaurant
		return t
	}
	log.Println(user.GetCuisines()[0].GetCuisineType())
	log.Println(user.GetCostBracket()[0].GetCostType())
	primaryCuisine := &CuisineFilter{user.GetCuisines()[0].GetCuisineType()}
	costFilter := &CostFilter{user.GetCostBracket()[0].GetCostType()}
	and := &AndFilter{
		FirstCritaria:  primaryCuisine,
		SecondCritaria: costFilter,
	}
	temp := and.Filter(restaurants)
	featured := &FeaturedrFilter{}
	temp = featured.Filter(temp)
	return temp
}

type FeaturedPrimaryCuisineSecondryCostStrategy struct{}

func (pc *FeaturedPrimaryCuisineSecondryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimaryCuisine() || !user.HasSecondaryCost() {
		var t []*models.Restaurant
		return t
	}
	secondaryCost1 := &CostFilter{user.GetCostBracket()[1].GetCostType()}
	secondaryCost2 := &CostFilter{user.GetCostBracket()[2].GetCostType()}
	or := &OrFilter{
		FirstCritaria:  secondaryCost1,
		SecondCritaria: secondaryCost2,
	}
	primeryFilter := &CuisineFilter{user.GetCuisines()[0].GetCuisineType()}
	and := &AndFilter{
		FirstCritaria:  or,
		SecondCritaria: primeryFilter,
	}
	temp := and.Filter(restaurants)
	featured := &FeaturedrFilter{}
	return featured.Filter(temp)
}

type FeaturedSecondaryCuisinePrimeryCostStrategy struct{}

func (pc *FeaturedSecondaryCuisinePrimeryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimryCost() || !user.HasSecondaryCuisine() {
		var t []*models.Restaurant
		return t
	}
	secondaryCuisine1 := &CuisineFilter{user.GetCuisines()[1].GetCuisineType()}
	secondaryCuisine2 := &CuisineFilter{user.GetCuisines()[2].GetCuisineType()}
	or := &OrFilter{
		FirstCritaria:  secondaryCuisine1,
		SecondCritaria: secondaryCuisine2,
	}
	costFilter := &CostFilter{user.GetCostBracket()[0].GetCostType()}
	and := &AndFilter{
		FirstCritaria:  or,
		SecondCritaria: costFilter,
	}
	temp := and.Filter(restaurants)
	featured := &FeaturedrFilter{}
	return featured.Filter(temp)
}

type PrimaryCuisinePrimaryCostStrategy struct{}

func (pc *PrimaryCuisinePrimaryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimaryCuisine() || !user.HasPrimryCost() {
		var t []*models.Restaurant
		return t
	}
	primaryCuisine := &CuisineFilter{user.GetCuisines()[0].GetCuisineType()}
	costFilter := &CostFilter{user.GetCostBracket()[0].GetCostType()}
	and := &AndFilter{
		FirstCritaria:  primaryCuisine,
		SecondCritaria: costFilter,
	}
	temp := and.Filter(restaurants)
	rating := &GreaterRating{4}
	return rating.Filter(temp)
}

type PrimaryCuisineSecondryCostStrategy struct{}

func (pc *PrimaryCuisineSecondryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimaryCuisine() || !user.HasSecondaryCost() {
		var t []*models.Restaurant
		return t
	}
	secondaryCost1 := &CostFilter{user.GetCostBracket()[1].GetCostType()}
	secondaryCost2 := &CostFilter{user.GetCostBracket()[2].GetCostType()}
	or := &OrFilter{
		FirstCritaria:  secondaryCost1,
		SecondCritaria: secondaryCost2,
	}
	primeryFilter := &CuisineFilter{user.GetCuisines()[0].GetCuisineType()}
	and := &AndFilter{
		FirstCritaria:  or,
		SecondCritaria: primeryFilter,
	}
	temp := and.Filter(restaurants)
	rating := &GreaterRating{4.5}
	return rating.Filter(temp)
}

type SecondaryCuisinePrimaryCostStrategy struct{}

func (pc *SecondaryCuisinePrimaryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasSecondaryCuisine() || !user.HasPrimryCost() {
		var t []*models.Restaurant
		return t
	}
	secondaryCuisine1 := &CuisineFilter{user.GetCuisines()[1].GetCuisineType()}
	secondaryCuisine2 := &CuisineFilter{user.GetCuisines()[2].GetCuisineType()}
	or := &OrFilter{
		FirstCritaria:  secondaryCuisine1,
		SecondCritaria: secondaryCuisine2,
	}
	costFilter := &CostFilter{user.GetCostBracket()[0].GetCostType()}
	and := &AndFilter{
		FirstCritaria:  or,
		SecondCritaria: costFilter,
	}
	temp := and.Filter(restaurants)
	rating := &GreaterRating{4.5}
	return rating.Filter(temp)
}

// Less rating
type LowPrimaryCuisinePrimaryCostStrategy struct{}

func (pc *LowPrimaryCuisinePrimaryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimaryCuisine() || !user.HasPrimryCost() {
		var t []*models.Restaurant
		return t
	}
	primaryCuisine := &CuisineFilter{user.GetCuisines()[0].GetCuisineType()}
	costFilter := &CostFilter{user.GetCostBracket()[0].GetCostType()}
	and := &AndFilter{
		FirstCritaria:  primaryCuisine,
		SecondCritaria: costFilter,
	}
	temp := and.Filter(restaurants)
	rating := &LessRating{4}
	return rating.Filter(temp)
}

type LowPrimaryCuisineSecondryCostStrategy struct{}

func (pc *LowPrimaryCuisineSecondryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasPrimaryCuisine() || !user.HasSecondaryCost() {
		var t []*models.Restaurant
		return t
	}
	secondaryCost1 := &CostFilter{user.GetCostBracket()[1].GetCostType()}
	secondaryCost2 := &CostFilter{user.GetCostBracket()[2].GetCostType()}
	or := &OrFilter{
		FirstCritaria:  secondaryCost1,
		SecondCritaria: secondaryCost2,
	}
	primeryFilter := &CuisineFilter{user.GetCuisines()[0].GetCuisineType()}
	and := &AndFilter{
		FirstCritaria:  or,
		SecondCritaria: primeryFilter,
	}
	temp := and.Filter(restaurants)
	rating := &LessRating{4.5}
	return rating.Filter(temp)
}

type LowSecondaryCuisinePrimaryCostStrategy struct{}

func (pc *LowSecondaryCuisinePrimaryCostStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {
	if !user.HasSecondaryCuisine() || !user.HasPrimryCost() {
		var t []*models.Restaurant
		return t
	}
	secondaryCuisine1 := &CuisineFilter{user.GetCuisines()[1].GetCuisineType()}
	secondaryCuisine2 := &CuisineFilter{user.GetCuisines()[2].GetCuisineType()}
	or := &OrFilter{
		FirstCritaria:  secondaryCuisine1,
		SecondCritaria: secondaryCuisine2,
	}
	costFilter := &CostFilter{user.GetCostBracket()[0].GetCostType()}
	and := &AndFilter{
		FirstCritaria:  or,
		SecondCritaria: costFilter,
	}
	temp := and.Filter(restaurants)
	rating := &LessRating{4.5}
	return rating.Filter(temp)
}

type NewltRestaurantStrategy struct{}

func (pc *NewltRestaurantStrategy) Filter(user *models.User, restaurants []*models.Restaurant) []*models.Restaurant {

	newlyRest := NewlyFilter{}
	return newlyRest.Filter(restaurants)
}
