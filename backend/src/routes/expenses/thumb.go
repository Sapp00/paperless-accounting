package expenses

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *ExpenseRouter) GetExpenseThumb(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "GET")
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not defined")
		return
	}

	idI, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	img, err := r.dm.GetThumb(idI)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", strconv.Itoa(len(img)))

	if _, err := c.Writer.Write(img); err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}
