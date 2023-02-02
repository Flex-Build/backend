// Package logoinit provides method to Init loging config
package logoinit

import (
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/sirupsen/logrus"
)

func Init() {
	logrusEntry := logrus.New().WithFields(logrus.Fields{})
	logo.SetInstance(*logrusEntry)
}
