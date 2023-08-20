package payments

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *PaymentRouter) DeletePayment(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "DELETE OPTIONS")

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

	err = r.dm.DeletePayment(idI)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	ret := result{Status: "ok"}

	c.JSON(http.StatusOK, ret)
}
