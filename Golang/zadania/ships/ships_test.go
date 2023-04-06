package ships

import "fmt"

func ExamplePoint_Add() {
	p := Point{1, 2}
	a := Point{3, 4}
	p = p.Add(a)
	fmt.Println(p.X, p.Y)
	// Output:
	// 4 6
}
func ExamplePoint_Add_Negative() {
	p := Point{1, 2}
	a := Point{-3, -4}
	p = p.Add(a)
	fmt.Println(p.X, p.Y) //println
	// Output:
	// -2 -2
}
func ExampleShip_MoveToNegative() {
	s := Ship{{1, 2}, {2, 2}, {3, 2}}
	s = s.MoveTo(Point{-4, -4})
	for _, v := range s {
		fmt.Println(v.X, v.Y)
	}
	// Output:
	// -4 -4
	// -3 -4
	// -2 -4
}
func ExampleShip_MoveTo() {
	s := Ship{{1, 2}, {2, 2}, {3, 2}}
	s = s.MoveTo(Point{4, 4})
	for _, v := range s {
		fmt.Println(v.X, v.Y) //println()
	}
	// Output:
	// 4 4
	// 5 4
	// 6 4
}
