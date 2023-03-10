package bootstrap

import (
	"Portfolio_Nodes/app"
	"Portfolio_Nodes/delivery/rest_delivery"
	"Portfolio_Nodes/interactors"
	"Portfolio_Nodes/repos"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"os"
)

func InitRest() error {
	if os.Getenv("APP_DEBUG") != "false" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())
	r.Use(rest_delivery.ErrorHandler)

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	// Initialization
	userRepo := repos.NewUserRepo(app.DB)
	tasksRepo := repos.NewTaskRepo(app.DB)
	//winnerRepo := repos.NewWinnerRepo(app.GormDB)
	//winnerRepo := repos.NewWinnerTRepo()

	//winnerInteractor := interactors.NewWinnerInteractor(winnerRepo)
	tasksIter := interactors.NewTaskInteractor(tasksRepo)
	userIter := interactors.NewUserIter(userRepo)
	// Registration

	rest_delivery.NewTaskDelivery(tasksIter).Route(r.Group("/task"))
	rest_delivery.NewUserDelivery(userIter).Router(r.Group("/user"))
	err = r.Run(":" + os.Getenv("REST_PORT"))
	if err != nil {
		return err
	}

	return nil
}
