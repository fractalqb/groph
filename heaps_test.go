package groph

import (
	"fmt"
)

func output(p []rune) {
	if len(p) > 0 {
		fmt.Printf("%c", p[0])
		for _, c := range p[1:] {
			fmt.Printf(", %c", c)
		}
	}
	fmt.Println()
}

func ExampleHeapsAlgo() {
	data := []rune{'a', 'b', 'c', 'd'}
	c := make([]int, len(data))
	for i := 0; i < len(c); i++ {
		c[i] = 0
	}
	output(data)
	i := 0
	for i < len(data) {
		if c[i] < i {
			if (i & 1) == 0 {
				data[0], data[i] = data[i], data[0]
			} else {
				data[c[i]], data[i] = data[i], data[c[i]]
			}
			output(data)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
}
