package env

import "os"

var appengineEnv AppengineEnv

func init() {
	appengineEnv.Port = os.Getenv("PORT")
	if appengineEnv.Port == "" {
		appengineEnv.Port = "8080"
	}
	appengineEnv.ProjectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	appengineEnv.Service = os.Getenv("GAE_SERVICE")
	appengineEnv.Version = os.Getenv("GAE_VERSION")
}

type AppengineEnv struct {
	Port      string
	ProjectID string
	Service   string
	Version   string
}

func Appengine() AppengineEnv {
	return appengineEnv
}
