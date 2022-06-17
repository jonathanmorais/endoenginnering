package main

// Importing fmt
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/t1d-api/controllers/calculate"
	"github.com/jonathanmorais/endoenginnering/t1d-api/controllers/healthcheck"
	"github.com/jonathanmorais/endoenginnering/t1d-api/pkg/memory_cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

// Calling main
func main() {

	// Logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	// use memory cache package
	c := memory_cache.GetInstance()
	// Readiness Probe Mock Config
	probeTimeRaw := os.Getenv("READINESS_PROBE_MOCK_TIME_IN_SECONDS")
	if probeTimeRaw == "" {
		probeTimeRaw = "5"
	}
	probeTime, err := strconv.ParseUint(probeTimeRaw, 10, 64)
	if err != nil {
		fmt.Println("Environment variable READINESS_PROBE_MOCK_TIME_IN_SECONDS conversion error", err)
	}
	c.Set("readiness.ok", "false", time.Duration(probeTime)*time.Second)

	// Creates a gin router with default middleware:
	router := gin.New()
	router.GET("/health", healthcheck.Ok)
	router.POST("/calculate", calculate.Calculate)

	router.Run()
}
