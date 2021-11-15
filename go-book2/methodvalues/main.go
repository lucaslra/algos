package main

import (
	"fmt"
	"time"
)

type Rocket struct {
	Name string
	Fuel int
}

func (r *Rocket) Launch() {
	r.Fuel = 0

	fmt.Println("Launching", r.Name)
}

func main() {
	r := Rocket{Name: "Apollo", Fuel: 100}

	time.AfterFunc(time.Second*5, r.Launch)

	time.Sleep(time.Second * 6)
}
