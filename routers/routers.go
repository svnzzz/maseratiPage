package routers

import (
	"maserati/luongo/entities/optionals"
	"maserati/luongo/entities/vehicle"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/vehicles", vehicle.GetVehicles)
	r.GET("/optionals", optionals.GetVehicleWithOptionals)
	return r
}
