package main

import (
	"confetti-framework/app/console"
	"confetti-framework/bootstrap"
	"os"
)

func main() {
	/*
	   |--------------------------------------------------------------------------
	   | Turn On The Lights
	   |--------------------------------------------------------------------------
	   |
	   |
	   |
	*/
	app := bootstrap.NewAppFromBoot()

	/*
	   |--------------------------------------------------------------------------
	   | Run The Application
	   |--------------------------------------------------------------------------
	   |
	   | Once we have the application, we can handle the command
	   | through the kernel.
	   |
	*/
	kernel := console.NewKernel(app)
	code := kernel.Handle()
	os.Exit(int(code))
}
