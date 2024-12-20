package tasks

import (
	"time"

	router "go-boilerplate-app/pkg/api/routers"
	"go-boilerplate-app/pkg/config"
	"go-boilerplate-app/pkg/logger"

	"github.com/spf13/cobra"
)

func SimpleServiceTask(cmd *cobra.Command, args []string) {
	now := time.Now()
	logger.Info("Time stamp is %d", now.UnixNano())
	router.InitSimpleAPIRouter()
	router.SimpleAPIRouter().Start(config.HostFlag, config.PortFlag)
}
