package main

import "main/routes"

//----------
// Handlers
//----------

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
