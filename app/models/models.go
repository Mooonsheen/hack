package models

// type CurrentPlanet struct {
// 	Name string
// 	PlanetTo []*Node
// }

// type CurrentPlanet struct {
// 	Planet   Planet `json:""`
// 	PlanetTo map[*Planet]int
// }

type Universe struct {
	Name       string          `json:"name"`
	RoundName  string          `json:"roundName"`
	RoundEndIn int             `json:"roundEndIn"`
	Ship       Ship            `json:"ship"`
	Universe   [][]interface{} `json:"universe"`
}

type Planet struct {
	NameTo string
	Oil    int
}

type CurrentPlanet struct {
	Name          string     `json:"name"`
	PlanetGarbage []*Garbage `json:"garbage"`
}

type Ship struct {
	FuelUsed    int           `json:"fuelUsed"`
	Planet      CurrentPlanet `json:"planet"`
	CapacityX   int           `json:"capacityX"`
	CapacityY   int           `json:"capacityY"`
	ShipGarbage []*Garbage    `json:"garbage"`
}

type Garbage struct {
	Name         string  `json:"name"`
	GarbageShape [][]int `json:"garbage"`
}
