package handler

// using request payload considering more fields can be added in future
type FuelCalculationRequest struct {
	CrewSize int `json:"crewSize" binding:"required,min=1"`
}

type FuelCalculateResponse struct {
	FuelCost float64 `json:"fuelCost"`
}
