package envinit

import "github.com/Flexi-Build/backend/pkg/envconfig"

func Init() {
	envconfig.InitEnvVars()
}
