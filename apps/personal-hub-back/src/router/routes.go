package router

import (
	"github.com/gin-gonic/gin"
)

func serveRoutes(router *gin.Engine) {

	investments := router.Group("/finance/investments")
	{
		investments.GET("/")                            //Returns all types of investiments, divided by stocks/non-stocks
		investments.GET("/stocks")                      //Returns all stocks types invested into
		investments.GET("stocks/positions")             //Returns all positions on stocks
		investments.GET("stocks/positions/:positionId") //Returns the positions on a given stock
		investments.POST("stocks/position")             //Register a position on a stock
	}

}
