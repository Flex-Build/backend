package main

import (
	"github.com/Flexi-Build/backend/app/stage/appinit"
	"github.com/Flexi-Build/backend/app/stage/apprun"
)

func main() {
	appinit.Init()
	apprun.Run()
}
