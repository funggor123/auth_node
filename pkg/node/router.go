package node

import (
	common "../common"
	"github.com/gin-gonic/gin"
)

func setGin() {
	gin.SetMode(GIN_MODE)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func StartRouter() {

	setGin()
	engine := gin.New()
	engine.Use(CORSMiddleware())

	route_register(engine, getNode())
	engine.Run(":" + common.GetConfiger().Configs.Port)
}

func route_register(router *gin.Engine, node *Node) {

	// v1 //
	v1 := router.Group("v1")
	{
		api := v1.Group("api")
		{
			api.PUT("/player", node.CreatePlayer)
			start := api.Group("start")
			{
				start.GET("/token", node.GetAgentToken)
			}
			wallet := api.Group("/wallet")
			{
				wallet.POST("/deposit", node.Deposit)
				wallet.POST("/withdraw", node.WithDraw)
				wallet.POST("/balance", node.GetBalance)
			}
			tran := api.Group("/transaction")
			{
				tran.GET("/get", node.GetTransactionByPID)
			}
		}
		api.POST("/gamerecords", node.GetGameRecordByPID)
		api.GET("/gamestrings", node.GetGameString)
		api.POST("/gamelaunch", node.LaunchGame)
	}

}
