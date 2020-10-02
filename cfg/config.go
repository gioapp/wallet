package cfg

// Conf is the configuration for accessing rpc endpoint
type Conf struct {
	Username, Password string
}

// configurations for wallet
var (
	//File = filepath.Join(Dir, "conf.json")
	//Web is a subfolder because otherwise the config above would be served by the http.Dir webserver
	//TSL = Dir + "/tsl/"
	//Web         = "/www/"
	Credentials = Conf{}
	Initial     bool
)
