package payments

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *PaymentRouter) GetPaymentsBetween(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "GET")
	from := c.DefaultQuery("from", "1940-01-01")
	to := c.DefaultQuery("to", "2036-12-12")

	fromT, err := time.Parse("2006-01-02", from)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	toT, err := time.Parse("2006-01-02", to)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	docs, err := r.dm.GetPaymentsBetween(fromT, toT)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, docs)
}
