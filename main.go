package main

import (
	"github.com/yaronsumel/geo-coding-test/src/dataHandling"
	"github.com/yaronsumel/geo-coding-test/src/dataHandling/handlers/csv"
	"github.com/yaronsumel/geo-coding-test/src/list"
	"github.com/yaronsumel/geo-coding-test/src/place"
	"log"
)

// our target point
var targetPoint = place.Place{Lat: 51.925146, Lon: 4.478617}
// places list to populate places
var placesList = list.PlacesList{}

// get things ready
func init() {
	// create new csv handler
	csvHandler, err := csv.NewCsvHandler("data.csv")
	if err != nil {
		log.Panicln(err)
	}
	// register handler
	dataHandling.RegHandler("csv", csvHandler)
}

func main() {

	// get handler
	handler, err := dataHandling.GetHandler("csv")
	if err != nil {
		log.Panicln(err)
	}

	// p will hold our place
	var p = &place.Place{}

	// iterate over results
	// will stops when EOF reached
	for handler.Next(p) {
		x := &place.Place{}
		// copy content of p into x
		*x = *p
		// calc distance
		x.CalcGreatCircleDistance(&targetPoint)
		// append x to list
		placesList.Append(x)
	}

	log.Printf("Closest Places")
	// range over the topFive and print them
	for _, v := range placesList.TopFive() {
		log.Printf("Place ID %v is %v Km away", v.Id, v.Distance)
	}

	log.Printf("Furthest Places")
	// range over the bottomFive and print them
	for _, v := range placesList.BottomFive() {
		log.Printf("Place ID %v is %v Km away", v.Id, v.Distance)
	}

	log.Println("Done!")
}
