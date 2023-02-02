// Package appinit provides method to Init all stages of app
package appinit

import (
	"github.com/Flexi-Build/backend/app/stage/appinit/dbconinit"
	"github.com/Flexi-Build/backend/app/stage/appinit/dbmigrate"
	"github.com/Flexi-Build/backend/app/stage/appinit/envinit"
	"github.com/Flexi-Build/backend/app/stage/appinit/logoinit"
)

func Init() {
	envinit.Init()
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
