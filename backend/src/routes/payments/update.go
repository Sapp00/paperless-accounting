package payments

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type updatePayment struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

type result struct {
	Status string `json:"status"`
}

func (r *PaymentRouter) UpdatePayment(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "POST")

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

	var updateData updatePayment
	err = c.BindJSON(&updateData)
	if err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, "Update Parameters not specified")
			return
		}
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var ptime *time.Time
	// if there is a date, convert it
	if updateData.Date != "" {
		time, err := time.Parse("2006-01-02", updateData.Date)
		if err != nil {
			log.Printf("Error occured: %s", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		ptime = &time
	}

	var value *float64
	if updateData.Value != "" {
		v, err := strconv.ParseFloat(updateData.Value, 64)
		if err != nil {
			log.Printf("Error occured: %s", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		value = &v
	}

	err = r.dm.UpdatePayment(idI, ptime, value)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	ret := result{Status: "ok"}

	c.JSON(http.StatusOK, ret)
}
