package main

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune('A')
		} else {
			z01.PrintRune('B')
		}
	}
	if y > 1 {
		for l := 1; l <= y-2; l++ {
			z01.PrintRune('\n')
			for k := 1; k <= x; k++ {
				if k == 1 || k == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
		for j := 1; j <= x; j++ {
			if j == 1 || j == x {
				z01.PrintRune('C')
			} else {
				z01.PrintRune('B')
			}
		}
	}
}

func main() {
	QuadC(1, 5)
	z01.PrintRune('\n')
}
