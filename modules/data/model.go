package data

import "time"

type Exoplanet struct {
	Id          int        `json:"id" `
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	Distance    int        `json:"distance" binding:"required,min=10,max=1000"`
	Radius      float64    `json:"radius" binding:"required,min=0.1,max=10"`
	Mass        float64    `json:"mass,omitempty" binding:"mass"`
	Type        string     `json:"type" binding:"exoplanetType"`
	FoundAt     *time.Time `json:"-"` // no need to send as a response these are for internal use
	UpdatedAt   *time.Time `json:"-"` // internal use
}

type Exoplanets []*Exoplanet

// will store the data in map, it will help to retrieve the data very fast
type ExoplanetMap map[int]*Exoplanet
