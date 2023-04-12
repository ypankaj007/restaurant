package models

type User struct {
	cuisines    []*CuisineTracking
	costBracket []*CostTracking
}

func NewUser(cuisines []*CuisineTracking, costBracket []*CostTracking) *User {
	return &User{
		cuisines:    cuisines,
		costBracket: costBracket,
	}
}

func (user *User) GetCuisines() []*CuisineTracking {
	return user.cuisines
}

func (user *User) GetCostBracket() []*CostTracking {
	return user.costBracket
}

func (user *User) SetCuisines(cuisines []*CuisineTracking) *User {
	user.cuisines = cuisines
	return user
}

func (user *User) SetCostBracket(costBracket []*CostTracking) *User {
	user.costBracket = costBracket
	return user
}

func (user *User) HasPrimryCost() bool {
	if user.costBracket == nil || len(user.costBracket) == 0 {
		return false
	}
	return true
}

func (user *User) HasPrimaryCuisine() bool {
	if user.cuisines == nil || len(user.cuisines) == 0 {
		return false
	}
	return true
}

func (user *User) HasSecondaryCost() bool {
	if user.costBracket == nil || len(user.costBracket) < 3 {
		return false
	}
	return true
}

func (user *User) HasSecondaryCuisine() bool {
	if user.cuisines == nil || len(user.cuisines) < 3 {
		return false
	}
	return true
}
