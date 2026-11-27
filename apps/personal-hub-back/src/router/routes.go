package router

import (
	"github.com/gin-gonic/gin"

	hobbies_controllers "github.com/personalhub/personalhub-back/src/controllers/hobbies"
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

	hobbies := router.Group("/hobbies")
	{
		cards := hobbies.Group("/cards")
		{
			cards.GET("/", hobbies_controllers.GetCards)         //Returns all cards
			cards.GET("/:id", hobbies_controllers.GetCard)       //Returns a single card
			cards.POST("/", hobbies_controllers.CreateCard)      //Registers a new card
			cards.PUT("/:id", hobbies_controllers.UpdateCard)    //Updates a card
			cards.DELETE("/:id", hobbies_controllers.DeleteCard) //Deletes a card
		}

		orders := hobbies.Group("/orders")
		{
			orders.GET("/", hobbies_controllers.GetOrders)         //Returns all orders
			orders.GET("/:id", hobbies_controllers.GetOrder)       //Returns a single order
			orders.POST("/", hobbies_controllers.CreateOrder)      //Registers a new order
			orders.PUT("/:id", hobbies_controllers.UpdateOrder)    //Updates an order
			orders.DELETE("/:id", hobbies_controllers.DeleteOrder) //Deletes an order
		}

		reimbursements := hobbies.Group("/reimbursements")
		{
			reimbursements.GET("/", hobbies_controllers.GetReimbursements)              //Returns all reimbursements
			reimbursements.GET("/:orderId", hobbies_controllers.GetReimbursement)       //Returns a single reimbursement
			reimbursements.POST("/", hobbies_controllers.CreateReimbursement)           //Registers a new reimbursement
			reimbursements.PUT("/:orderId", hobbies_controllers.UpdateReimbursement)    //Updates a reimbursement
			reimbursements.DELETE("/:orderId", hobbies_controllers.DeleteReimbursement) //Deletes a reimbursement
		}
	}

}
