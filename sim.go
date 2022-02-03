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

		case 1:
			g := .002 // [pixel/frame^2]

			vx[t][n] = 2.6            // [pixel/frame]
			vy[t][n] = g * float64(t) // [pixel/frame]

			x[t][n] = 0.0 + vx[t][n]*float64(t)   // [pixel]
			y[t][n] = 140.0 + vy[t][n]*float64(t) // [pixel]
		case 2:
			g := .003 // [pixel/frame^2]
			switch t {
			case 0:
				vx[t][n] = 1.8 // [pixel/frame]
				vy[t][n] = -.7 // [pixel/frame]

				x[t][n] = 0.0   // [pixel]
				y[t][n] = 400.0 // [pixel]
			default:
				vx[t][n] = vx[t-1][n]     // [pixel/frame]
				vy[t][n] = vy[t-1][n] + g // [pixel/frame]

				x[t][n] = x[t-1][n] + vx[t-1][n] // [pixel]
				y[t][n] = y[t-1][n] + vy[t-1][n] // [pixel]
			}
		}
	}
}
