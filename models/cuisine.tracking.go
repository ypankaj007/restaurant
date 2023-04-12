package models

type CuisineTracking struct {
	cuisineType Cuisine // type
	noOfOrders  int
}

func NewCuisineTracking(cuisineType Cuisine, noOfOrders int) *CuisineTracking {
	return &CuisineTracking{
		cuisineType: cuisineType,
		noOfOrders:  noOfOrders,
	}
}

func (cuisinetracking *CuisineTracking) GetCuisineType() Cuisine {
	return cuisinetracking.cuisineType
}

func (cuisinetracking *CuisineTracking) GetNoOfOrders() int {
	return cuisinetracking.noOfOrders
}

func (cuisinetracking *CuisineTracking) SetCuisineType(cuisineType Cuisine) *CuisineTracking {
	cuisinetracking.cuisineType = cuisineType
	return cuisinetracking
}

func (cuisinetracking *CuisineTracking) SetNoOfOrders(noOfOrders int) *CuisineTracking {
	cuisinetracking.noOfOrders = noOfOrders
	return cuisinetracking
}
