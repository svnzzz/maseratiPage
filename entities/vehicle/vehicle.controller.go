package vehicle

import (
	"maserati/luongo/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET /vehicles
//GET vechicles/:id/optionals

func GetVehicles(c *gin.Context) {
	rows, err := initializers.DB.Query("SELECT * FROM TModelli")
	if err != nil {
		return
	}
	defer rows.Close()

	var vehicle []Vehicle

	for rows.Next() {
		var veh Vehicle
		if err := rows.Scan(&veh.ID, &veh.Model, &veh.BasePrice, &veh.BackgroundImage); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		vehicle = append(vehicle, veh)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(vehicle) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func GetVehiclesOptionals(c *gin.Context) {

}
