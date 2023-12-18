package complibgo

func AreaPolygon(polygon []Point) float64 {
	sum := 0.0
	n := len(polygon)

	for i := 0; i < n-1; i++ {
		sum += polygon[i].x*polygon[i+1].y - polygon[i].y*polygon[i+1].x
	}
	sum += polygon[n-1].x*polygon[0].y - polygon[n-1].y*polygon[0].x
	return sum / 2.0
}
