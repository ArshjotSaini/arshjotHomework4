package utils

import (
	"image/color"
	"strconv"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	simpleMap "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
)

const IMAGEPATH = "./assets/nameofimage.png"

var LocationCache map[string]*geo.Location
var JobCounts map[string]int

func Mapinit() {
	if !IsFileExist(IMAGEPATH) {
		LocationCache = make(map[string]*geo.Location)
		JobCounts = make(map[string]int)
		details := FetchCompanyData(db_conn)
		var loc *geo.Location
		for _, job := range details {
			loc = FindLocation(job.Location)
			ShowMap([][]string{{strconv.FormatFloat(loc.Lat, 'f', -1, 64)}, {strconv.FormatFloat(loc.Lng, 'f', -1, 64)}}, loc)
		}
	}
}

func FindLocation(city string) *geo.Location {
	gepLookup := openstreetmap.Geocoder()
	locationData, err := gepLookup.Geocode(city)
	ThrowError(err)
	return locationData
}

func ShowMap(data [][]string, loc *geo.Location) {
	processData(data, loc)
	context := simpleMap.NewContext()
	context.SetSize(1200, 1200)
	context.SetZoom(6)
	for city, numJobs := range JobCounts {
		cityLoc := LocationCache[city]
		pinColor := gotColor(numJobs)
		context.AddObject(
			simpleMap.NewMarker(
				s2.LatLngFromDegrees(cityLoc.Lat, cityLoc.Lng),
				pinColor,
				16,
			),
		)
	}

	context.SetCenter(s2.LatLngFromDegrees(loc.Lat, loc.Lng))
	image, err := context.Render()
	ThrowError(err)
	if err := gg.SavePNG(IMAGEPATH, image); err != nil {
		panic(err)
	}

}

func gotColor(numberOfjobs int) color.RGBA {
	if numberOfjobs > 75 {
		return color.RGBA{0, 0xff, 0, 0xff}
	}
	if numberOfjobs > 50 {
		return color.RGBA{
			R: 0,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		}
	}
	if numberOfjobs > 25 {
		return color.RGBA{
			R: 0,
			G: 0,
			B: 0xff,
			A: 0xff,
		}
	}
	if numberOfjobs > 10 {
		return color.RGBA{
			R: 0xff,
			G: 0,
			B: 0xff,
			A: 0xff,
		}
	}
	if numberOfjobs > 5 {
		return color.RGBA{
			R: 0xff,
			G: 000,
			B: 000,
			A: 0xff,
		}

	}
	if numberOfjobs > 1 {
		return color.RGBA{
			R: 0x10,
			G: 0x10,
			B: 0x10,
			A: 0xff,
		}
	}
	return color.RGBA{}
}

func processData(data [][]string, defaultLocation *geo.Location) {
	for rowNumber, job := range data {
		if rowNumber < 1 {
			continue
		}
		cityName := job[0]
		_, ok := LocationCache[cityName]
		if ok {
			JobCounts[cityName]++
		} else {
			loc := FindLocation(cityName)
			if loc == nil {
				loc = defaultLocation
			}
			LocationCache[cityName] = loc
			JobCounts[cityName] = 1
		}
	}
}
