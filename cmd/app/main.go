package main

import (
	"binance-order-matcher/config"
	"binance-order-matcher/internal/controller/http"
	"binance-order-matcher/internal/service/repo"
	"binance-order-matcher/pkg/httpserver"
	"binance-order-matcher/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	conf, err := config.Init()
	if err != nil {
		log.Panic("while getting config", err)
		os.Exit(1)
	}
	l := logger.New(conf.Log.Level)

	m, err := migrate.New("file://migrations", conf.Database.URL)
	if err != nil {
		l.Error("while connecting to database. ", err)
		os.Exit(1)
	}
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			l.Info("migrate: no change")
		} else {
			l.Error("while migrating database. ", err)
			os.Exit(1)
		}
	}
	u, err := url.Parse(conf.Database.URL)
	if err != nil {
		l.Error("malformed database URL", err)
		os.Exit(1)
	}
	db, err := gorm.Open(sqlite.Open(u.Host), &gorm.Config{})
	if err != nil {
		l.Error("while connecting to database. ", err)
		os.Exit(1)
	}
	userRepo := repo.NewUserRepo(db)
	// HTTP Server
	handler := gin.New()
	http.NewRouter(handler, l, userRepo)
	httpServer := httpserver.New(handler, httpserver.Port(conf.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
