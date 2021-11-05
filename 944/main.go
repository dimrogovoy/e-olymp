package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, z float64
}

type Triangle struct {
	ab, bc, ca float64
}

func (t *Triangle) distance(x Point, y Point, z Point) {
	t.ab = math.Sqrt((y.x-x.x)*(y.x-x.x) + (y.y-x.y)*(y.y-x.y) + (y.z-x.z)*(y.z-x.z))
	t.bc = math.Sqrt((z.x-y.x)*(z.x-y.x) + (z.y-y.y)*(z.y-y.y) + (z.z-y.z)*(z.z-y.z))
	t.ca = math.Sqrt((x.x-z.x)*(x.x-z.x) + (x.y-z.y)*(x.y-z.y) + (x.z-z.z)*(x.z-z.z))
}
func (t *Triangle) area() float64 {
	var p float64
	p = (t.ab + t.bc + t.ca) / 2
	return math.Sqrt(p * (p - t.ab) * (p - t.bc) * (p - t.ca))

}

func main() {
	//var a,b,c,d Point
	//fmt.Scanf("%f %f %f\n%f %f %f\n%f %f %f\n%f %f %f",&a.x,&a.y,&a.z,&b.x,&b.y,&b.z,&c.x,&c.y,&c.z,&d.x,&d.y,&d.z)
	//var c1,c2, c3,c4 Triangle
	//	c1.distance(a,b,c)
	//	c2.distance(a,b,d)
	//	c3.distance(a,c,d)
	//	c4.distance(b,c,d)

	arr := make([]Point, 4)
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%f %f %f", &arr[i].x, &arr[i].y, &arr[i].z)
	}

	var c1, c2, c3, c4 Triangle
	c1.distance(arr[0], arr[1], arr[2])
	c2.distance(arr[0], arr[1], arr[3])
	c3.distance(arr[0], arr[2], arr[3])
	c4.distance(arr[1], arr[2], arr[3])

	fmt.Printf("%.1f\n", c1.area()+c2.area()+c3.area()+c4.area())

}
