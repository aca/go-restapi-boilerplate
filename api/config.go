package api

import (
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// DB
	ConfigDBURN    = "db.urn"
	ConfigDBDriver = "db.driver"

	// Log ("json" || "")
	ConfigLogFormat = "log"

	// http
	ConfigHTTPAddr = "http"

	// certificate
	ConfigTLSCert = "tls.crt"
	ConfigTLSKey  = "tls.key"
)

func buildFlags() *pflag.FlagSet {
	flags := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	flags.String(ConfigLogFormat, "", "log format")
	flags.String(ConfigDBURN, ":memory:?_foreign_keys=on", "dburn")
	flags.String(ConfigDBDriver, "sqlite3", "db driver(mysql/sqlite3)")
	flags.String(ConfigHTTPAddr, ":8080", "HTTP API listen address")

	flags.String(ConfigTLSCert, "", "tls certificate for api token")
	flags.String(ConfigTLSKey, "", "tls key for api token")

	return flags
}

func Configure(args []string) *viper.Viper {
	v := viper.New()

	// Setup command line flags
	flags := buildFlags()
	if err := flags.Parse(args); err != nil {
		panic(err)
	}

	// Configuration from flags
	if err := v.BindPFlags(flags); err != nil {
		panic(err)
	}

	v.SetEnvPrefix(os.Args[0])

	// Configuration from env
	v.AutomaticEnv()
	return v
}
