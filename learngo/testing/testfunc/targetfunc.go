package testfunc

import (
	"fmt"
	"math"
)

func triangleNormal(){
	var a,b int = 3,4
	c := int(math.Sqrt(float64(a*a +b*b)))
	fmt.Println(c)
}


func triangle(){
	var a,b int =3,4
	fmt.Println(CalcTriangle(a,b))
}

func CalcTriangle(a,b int) int{
	var c int
	c = int(math.Sqrt(float64(a*a +b*b)))
	return c
}


