package ols

//This function accepts two arrays (slices) a dependent and an independent variable. It returns the coefficients for a line of best fit given the Ordinary Least Squares Method as well as the R^2 value.

func reg(x, y []float64) (float64, float64, float64) {

	beta := (sum(multiply(dif(x), dif(y))) / sum(square(dif(x))))

	beta0 := float64(mean(y) - (beta * mean(x)))

	yres := make([]float64, 0)

	for i := range x {
		yres = append(yres, (y[i] - (x[i]*beta + beta0)))
	}

	r2 := 1 - sum(square(yres))/sum(square(dif(y)))

	return beta, beta0, r2

}

func mean(m []float64) float64 {

	tot := float64(0)

	for i := range m {
		tot += m[i]
	}

	return float64(tot / float64(len(m)))

}
func sum(m []float64) (tot float64) {
	for i := range m {
		tot += m[i]
	}
	return
}

func dif(x []float64) (t []float64) {
	t = make([]float64, 0)

	for i := range x {
		t = append(t, (x[i] - mean(x)))
	}
	return

}

func multiply(x []float64, y []float64) (t []float64) {
	t = make([]float64, 0)
	for i := range x {
		t = append(t, x[i]*y[i])
	}
	return

}

func square(m []float64) (sq []float64) {
	for i := range m {
		sq = append(sq, m[i]*m[i])
	}
	return
}
