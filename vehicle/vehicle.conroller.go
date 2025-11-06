package vehicle

import (
	"maserati/luongo/initializers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET /vehicles
//GET vechicles/{model}/optionals

func GetVehicles(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rows, err := initializers.DB.Query("SELECT * FROM TModelli WHERE ModelloID = @p1", idInt)
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
