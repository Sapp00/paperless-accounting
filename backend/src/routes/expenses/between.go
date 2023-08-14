package expenses

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *ExpenseRouter) GetExpensesBetween(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "GET")
	fromT := c.DefaultQuery("from", "1940-01-01")
	toT := c.DefaultQuery("to", "2036-12-12")

	docs, err := r.dm.GetExpensesBetween(fromT, toT)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, docs)
}
