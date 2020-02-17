package main

import (
	"fmt"
	"email_tflite_go/app"
)

func main() {

	a := app.CreateApplication()
	fmt.Println("Application ", a)
}
