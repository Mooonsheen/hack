package main

import (
	"log"
	"test/models"
	"test/responser"
)

// var Universe map[string][]models.Planet

func main() {
	Universe := make(map[string][]models.Planet)
	uni := responser.GetCurrentPlanet(&responser.Request{
		Url: "https://datsedenspace.datsteam.dev/player/universe",
	})

	// for i, v := range uni.Universe {
	// 	log.Printf("Planet %d: %v", i, v)
	// }

	getMap(&Universe)

	log.Printf("Our map: %v", Universe)
}

func getMap(u *map[string][]models.Planet) {
	i := 0
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
				}
				log.Printf("%v", v)
			case int:
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
}
