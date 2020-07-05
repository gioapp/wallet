package cfg

import (
	"path/filepath"

	"github.com/gioapp/jorm/pkg/appdata"
)

// Conf is the configuration for accessing bitnodes endpoint
type Conf struct {
	Username, Password, CloudFlareAPI, CloudFlareEmail, CloudFlareAPIkey string
}

// configurations for jorm
var (
	Dir  = appdata.Dir("jorm", false)
	File = filepath.Join(Dir, "conf.json")
	// Web is a subfolder because otherwise the config above would be served by the http.Dir webserver
	TSL         = Dir + "/tsl/"
	Web         = "/www/"
	Credentials = Conf{}
	Initial     bool
)
