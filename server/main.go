package main

import (
	"log"

	"github.com/goellavish10/img-upload-go/app"

	_ "github.com/goellavish10/img-upload-go/docs"
)

// @title			The Image Uploader App
// @version		1.0
// @description	A Golang application to test Devops and Backend concepts
// @contact.name	Lavish Goyal
// @license.name	MIT
// @host			localhost:8080
// @BasePath		/
func main() {
	err := app.SetupAppAndRun()

	if err != nil {
		log.Fatal(err)
	}

}
