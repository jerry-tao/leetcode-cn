type Solution struct {
	x float64
	y float64
	r float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		x_center, y_center, radius,
	}
}

func (this *Solution) RandPoint() []float64 {
	l := math.Sqrt(rand.Float64())*this.r
	a := math.Pi*2*rand.Float64()
	x := this.x + l*math.Sin(a)
	y := this.y + l*math.Cos(a)
	return []float64{x, y}
}
