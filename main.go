package payment_service

import (
	"github.com/PoojaSrinivasan18/payment-service/common"
	"github.com/PoojaSrinivasan18/payment-service/database"
	payment_service "github.com/PoojaSrinivasan18/payment-service/payment-service"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

func main() {
	err := common.ConfigSetup("configuration/dbconfig.yaml")
	if err != nil {
		log.Error("ConfigSetup failed")
		return
	}

	configuration := common.GetConfig()
	err = database.SetupDB(configuration)

	if err != nil {
		log.Error("SetupDB failed")
		return
	} else {
		log.Info("DB Setup Success")
	}

	router := gin.Default()
	router.GET("/api/getpaymentbyid", payment_service.GetPaymentById)
	router.POST("/api/makepayment", payment_service.MakePayment)
	router.DELETE("/api/deletepayment", payment_service.DeletePayment)

	//router.POST("/api/loadseeddata", inventory.SeedInventoryDetail)*/

	//:: Note: For local testing use below
	router.Run("localhost:3000")

	//:: For Docker use below
	//router.Run(":3000")
}
