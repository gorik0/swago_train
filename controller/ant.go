package controller

import (
	"github.com/gin-gonic/gin"
	"swago/models"
)



// GetCollonies godoc
//
//	@Summary		List Collonies population
//	@Description	get collonies
//	@Tags			ant
//	@Accept			text/plain
//	@Produce		json
//	@Success		200 	{object}		models.ColonyList
//	@Failure		400 	"ssd"
//	@Router			/ant [get]
func (c *Controller) GetCollonies(ctx *gin.Context) {

	var colList models.ColonyList



	for _, ant := range models.Ants {
		colList.PopulationList = append(colList.PopulationList, ant.Quantity)
	}
	ctx.Header("Content-Type","application/json")
	ctx.JSON(200,colList)
}
