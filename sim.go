package main

// Simulate motion state for the time t
func sim(t int) {
	for n := range x[t] {
		switch n {

		case 0:
			vx[t][n] = 2.6 // [pixel/frame]
			vy[t][n] = 0.0 // [pixel/frame]

			x[t][n] = 0.0 + vx[t][n]*float64(t)   // [pixel]
			y[t][n] = 540.0 + vy[t][n]*float64(t) // [pixel]

		case 9:
			a := .002 // [pixel/frame^2]

			vx[t][n] = 2.6            // [pixel/frame]
			vy[t][n] = a * float64(t) // [pixel/frame]

			x[t][n] = 0.0 + vx[t][n]*float64(t)   // [pixel]
			y[t][n] = 140.0 + vy[t][n]*float64(t) // [pixel]
		}
	}
}
