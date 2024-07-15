package gofakeit

import (
	"github.com/brianvoe/gofakeit/v7/data"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
	"strings"
)

// GeoShape will return a random slice of points (a shape) contained within the provided state shape
// if state or numOfPoints is default value (empty string or 0) an empty [][]float64 is returned
func GeoShape(state string, numOfPoints int) [][]float64 {
	return geoShape(GlobalFaker, state, numOfPoints)
}

// GeoShape will return a random slice of points (a shape) contained within the provided state shape
// if state or numOfPoints is default value (empty string or 0) an empty [][]float64 is returned
func (f *Faker) GeoShape(state string, numOfPoints int) [][]float64 {
	return geoShape(f, state, numOfPoints)
}

func geoShape(f *Faker, state string, numOfPoints int) [][]float64 {
	if state == "" || numOfPoints == 0 {
		return [][]float64{}
	}

	stateCoords, ok := data.StateGeoShapes[strings.ToLower(state)]
	if !ok {
		return nil
	}
	// Convert the state coordinates to orb.Polygon
	var polygons []orb.Polygon
	for _, poly := range stateCoords {
		var ring orb.Ring
		for _, point := range poly {
			ring = append(ring, orb.Point{point[0], point[1]})
		}
		polygons = append(polygons, orb.Polygon{ring})
	}

	// Find the bounding box of the state
	bound := polygons[0].Bound()
	for _, poly := range polygons[1:] {
		bound = bound.Union(poly.Bound())
	}
	var points []orb.Point
	var coords [][]float64
	for len(points) < numOfPoints {
		// Generate a random point within the bounding box
		x := f.Float64()*(bound.Max[0]-bound.Min[0]) + bound.Min[0]
		y := f.Float64()*(bound.Max[1]-bound.Min[1]) + bound.Min[1]
		point := orb.Point{x, y}
		// Check if the point is inside any of the polygons
		for _, poly := range polygons {
			if planar.PolygonContains(poly, point) {
				points = append(points, point)
				coords = append(coords, []float64{point[0], point[1]})
				break
			}
		}
	}
	return coords
}
