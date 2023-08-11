package expenses

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *ExpenseRouter) GetExpense(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
		c.JSON(http.StatusBadRequest, "ID not defined")
		return
	}
	year := c.Param("year")
	if year == "" {
		c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
		c.JSON(http.StatusBadRequest, "Year not defined")
		return
	}

	idI, err := strconv.Atoi(id)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
		c.JSON(http.StatusBadRequest, "Invalid ID")
		return
	}

	doc, err := r.dm.GetExpense(idI)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.JSON(http.StatusOK, doc)
}
