package expenses

import (
	"fmt"
	"log"
	"net/http"
	"sapp/paperless-accounting/paperless"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type updateExpense struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

func (r *ExpenseRouter) UpdateExpense(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT")
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

	var updateData updateExpense
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

	fmt.Printf("json: %v\n", updateData)

	var ptime *paperless.PaperlessTime
	// if there is a date, convert it
	if updateData.Date != "" {
		time, err := time.Parse("2006-01-02", updateData.Date)
		if err != nil {
			log.Printf("Error occured: %s", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		ptime = paperless.NewPaperlessTime(time)
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

	doc, err := r.dm.UpdateExpense(idI, ptime, value)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, doc)
}
