package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/dkr290/go-projects/gomonitoring/internal/channeldata"
	"github.com/dkr290/go-projects/gomonitoring/internal/driver"
	"github.com/pusher/pusher-http-go"
	"github.com/robfig/cron/v3"
)

// AppConfig holds application configuration
type AppConfig struct {
	DB            *driver.DB
	Session       *scs.SessionManager
	InProduction  bool
	Domain        string
	MonitorMap    map[int]cron.EntryID
	PreferenceMap map[string]string
	Scheduler     *cron.Cron
	WsClient      pusher.Client
	PusherSecret  string
	TemplateCache map[string]*template.Template
	MailQueue     chan channeldata.MailJob
	Version       string
	Identifier    string
}
