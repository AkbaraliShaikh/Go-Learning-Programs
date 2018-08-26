package main

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

func main() {

	var pixel [WIDTH][HEIGHT]int

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			pixel[x][y] = 0
		}
	}
}
