package complibgo

func AreaPolygon(polygon []Point) float64 {
	sum := 0.0
	n := len(polygon)

	for i := 0; i < n-1; i++ {
		sum += polygon[i].X*polygon[i+1].Y - polygon[i].Y*polygon[i+1].X
	}
	sum += polygon[n-1].X*polygon[0].Y - polygon[n-1].Y*polygon[0].X
	return sum / 2.0
}
