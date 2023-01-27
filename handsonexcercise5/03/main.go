package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "orange",
		},
		fourWheel: false,
	}

	t2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "jet black",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t1.color)
	fmt.Println(t2.color, t2.luxury)

}
