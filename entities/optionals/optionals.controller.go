package optionals

import (
	"fmt"
	"maserati/luongo/initializers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVehicleWithOptionals(c *gin.Context) {
	id := c.Query("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	rows, err := initializers.DB.Query("SELECT ModelloID, Nome, Prezzo, Modello, FileImmagine FROM viewModelliOptional WHERE ModelloID = @p1", intId)

	if err != nil {
		return
	}
	defer rows.Close()

	var optionals []Optional

	for rows.Next() {
		var opt Optional
		if err := rows.Scan(&opt.ModelID, &opt.Name, &opt.Price, &opt.Model, &opt.FileImage); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		optionals = append(optionals, opt)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(optionals) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	fmt.Print(optionals)
	c.JSON(http.StatusOK, optionals)

}
