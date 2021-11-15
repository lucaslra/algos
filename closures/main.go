package main

import (
	dataisolation "closures/pkgs/dataIsolation"
	"closures/pkgs/isolation"
	"closures/pkgs/middleware"
	"closures/pkgs/settear"
	"fmt"
	"net/http"
)

func main() {
	runIsolation()
	runSetTear()
	runIsolatedData()
	runMiddleware()

	// log.Fatal(http.ListenAndServe(":3000", nil))
}

func runSetTear() {
	defer settear.Setup()()
}

func runIsolatedData() {
	db := dataisolation.NewDatabase("TestName", "TestUrl")

	route := dataisolation.DatabaseRouteGen(db)
	http.HandleFunc("/db", middleware.TimedRoute(route))
}

func runMiddleware() {
	http.HandleFunc("/hello", middleware.TimedRoute(middleware.HelloWorldRoute))
}

func runIsolation() {
	gen := isolation.MakeFibonacciGen()

	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}
