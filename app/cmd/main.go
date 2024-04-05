package main

import (
	"fmt"
	"log"
	"test/app/models"
	"test/app/responser"
)

// var Universe map[string][]models.Planet

func main() {
	uni := responser.GetCurrentPlanet(&responser.Request{
		Url: "http://localhost:8081/universe", // https://datsedenspace.datsteam.dev/player/universe
	})

	// for i, v := range uni.Universe {
	// 	log.Printf("Planet %d: %v", i, v)
	// }

	uniMap := getMap(*uni)

	log.Printf("Planet %v:", uniMap)
}

func getMap(u models.Universe) *map[string][]models.Planet {
	Universe := make(map[string][]models.Planet)
	i := 0
	count := 0
	var planet1, planet2 string
	var neededOil int
	for _, item := range u.Universe {
		for _, value := range item {
			switch v := value.(type) {
			case string:
				var val string = value.(string)
				if i%2 == 0 {
					planet1 = val
				} else {
					planet2 = val
					i = 1
				}
				i++
				fmt.Printf("%v", v)
			case float64:
				count++
				var val int = int(value.(float64))
				neededOil = val
			default:
				count++
				var val int = value.(int)
				neededOil = val
			}
		}
		if _, ok := Universe[planet1]; !ok {
			Universe[planet1] = []models.Planet{{NameTo: planet2, Oil: neededOil}}
		} else {
			Universe[planet1] = append(Universe[planet1], models.Planet{NameTo: planet2, Oil: neededOil})
		}
	}
	log.Printf("count: %d", count)
	return &Universe
}
