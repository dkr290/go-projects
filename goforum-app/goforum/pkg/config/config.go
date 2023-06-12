package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

//store system wide configuration

type AppConfig struct {
	InfoLog   *log.Logger
	Session   *scs.SessionManager
	CSRFToken string
}
