package app

import (
	"Portfolio_Nodes/app/logs_hooks"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogs() error {
	//log.SetFormatter(&easy.Formatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	LogFormat:       "[%lvl%]: %time% - %msg%\n",
	//})

	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
	log.SetReportCaller(true)
	if os.Getenv("APP_DEBUG") == "true" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// hooks
	log.AddHook(logs_hooks.NewToFileHook())
	log.AddHook(logs_hooks.NewToFileErrorHook())

	return nil
}
