package main

import (
	"fmt"
	randomdata "github.com/Pallinder/go-randomdata"
	ellipse "github.com/bounaykut2/GoTests/ellipse"
)

func main() {
	fmt.Println(randomdata.City())
	//Initalise the Init function with value of A,B
	e := ellipse.Init{
		9, 2,
	}
	//this will give answer as 0.9749960430435691
	fmt.Println(e.GetEccentricity())
}
