//go:generate echo "oapi-codegen"
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=api --generate chi-server,spec -o api.gen.go ../spec/swagger.yaml

package api

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aca/go-restapi-boilerplate/ent"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type server struct {
	db         *ent.Client
	httpClient *http.Client
	v          *viper.Viper
	root       http.Handler
}

func NewServer(ctx context.Context, v *viper.Viper) (*server, error) {
	var err error

	s := &server{
		v: v,
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Logger.With().Timestamp().Caller().Logger()

	// global logger
	if v.GetString(ConfigLogFormat) == "json" {
		log.Logger = zerolog.New(os.Stderr)
	} else {
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// db
	s.db, err = ent.Open(v.GetString(ConfigDBDriver), v.GetString(ConfigDBURN))
	if err != nil {
		return nil, err
	}

	if err := s.db.Schema.Create(ctx); err != nil {
		return nil, err
	}

	// configure http client for global usage
	s.httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	// routers, middlewares
	r := chi.NewRouter()

	r.Handle("/metrics", promhttp.Handler())

	r.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(hlog.NewHandler(log.Logger))
		r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("")
		}))
		r.Use(hlog.RemoteAddrHandler("ip"))
		r.Use(hlog.UserAgentHandler("user_agent"))
		r.Use(hlog.RefererHandler("referer"))

		// If you are service is behind load balancer like nginx, you might want to
		// use X-Request-ID instead of injecting request id. You can do some thing
		// like this,
		// r.Use(hlog.CustomHeaderHandler("reqId", "X-Request-Id"))
		r.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
		r.Use(mwMetrics)
		r.Handle("/", HandlerFromMux(s, r))
	})

	// health check
	r.HandleFunc("/ping", ping)

	s.root = r

	return s, nil
}

// ping is handler responding to health-check request
func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.root.ServeHTTP(w, r)
}
