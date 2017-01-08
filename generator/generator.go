package main

import "fmt"

var G [256][256]string

func init() {
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			delta := y - x
			if delta > 128 {
				delta -= 256
			}

			if delta < -128 {
				delta += 256
			}

			if delta >= 0 {
				G[x][y] = repeat("+", delta)
			} else {
				G[x][y] = repeat("-", -delta)
			}
		}
	}

	iter := true
	for iter {
		iter = false
		for x := 0; x < 256; x++ {
			for n := 1; n < 40; n++ {
				for d := 1; d < 40; d++ {
					j := x
					y := 0
					for i := 0; i < 256; i++ {
						if j == 0 {
							break
						}

						j = (j - d + 256) & 255
						y = (y + n) & 255
					}
					if j == 0 {
						s := "[" + repeat("-", d) + ">" + repeat("+", n) + "<]>"
						if len(s) < len(G[x][y]) {
							G[x][y] = s
							iter = true
						}
					}

					j = x
					y = 0
					for i := 0; i < 256; i++ {
						if j == 0 {
							break
						}
						j = (j + d) & 255
						y = (y - n + 256) & 255
					}

					if j == 0 {
						s := "[" + repeat("+", d) + ">" + repeat("-", n) + "<]>"
						if len(s) < len(G[x][y]) {
							G[x][y] = s
							iter = true
						}
					}
				}
			}
		}

		for x := 0; x < 256; x++ {
			for y := 0; y < 256; y++ {
				for z := 0; z < 256; z++ {
					if len(G[x][z])+len(G[z][y]) < len(G[x][y]) {
						G[x][y] = G[x][z] + G[z][y]
						iter = true
					}
				}
			}
		}
	}
}

func generate(s string) {
	lastch := 0
	for _, c := range s {
		a := G[lastch][int(c)]
		b := G[0][int(c)]
		if len(a) <= len(b) {
			fmt.Print(a)
		} else {
			fmt.Print(">" + b)
		}
		fmt.Print(".")
		lastch = int(c)
	}
	fmt.Println("\n")
}

func repeat(s string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += s
	}
	return ret
}
