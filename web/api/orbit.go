package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrbit 获取用户围栏
// @Tags orbit
// @Summary 获取用户围栏
// @Description 获取用户围栏
// @Accept  json
// @Produce  json
// @Param 	Authorization 	header 	string 	true "With the bearer started JWT"
// @Param	offset	id		int	flase	"用户ID"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /orbit [get]
func GetOrbit(c *gin.Context) {
	ids := c.Query("id")
	id := 0
	var err error
	if ids == "" {
		id = c.GetInt("id")
	} else {
		id, err = strconv.Atoi(ids)
		if err != nil {
			retError(c, 12, err)
			return
		}
		if id == 0 {
			retError(c, 12, err)
			return
		}
	}

}
