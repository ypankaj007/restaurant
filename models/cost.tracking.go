package models

type CostTracking struct {
	costType   int // type
	noOfOrders int
}

func NewCostTracking(costType, noOfOrders int) *CostTracking {
	return &CostTracking{
		costType:   costType,
		noOfOrders: noOfOrders,
	}
}
func (costtracking *CostTracking) GetCostType() int {
	return costtracking.costType
}

func (costtracking *CostTracking) GetNoOfOrders() int {
	return costtracking.noOfOrders
}

func (costtracking *CostTracking) SetCostType(costType int) *CostTracking {
	costtracking.costType = costType
	return costtracking
}

func (costtracking *CostTracking) SetNoOfOrders(noOfOrders int) *CostTracking {
	costtracking.noOfOrders = noOfOrders
	return costtracking
}
