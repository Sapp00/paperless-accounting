package payments

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type createPayment struct {
	Date      string `json:"date"`
	Value     string `json:"value"`
	ExpenseID int    `json:"expense"`
}

func (r *PaymentRouter) CreatePayment(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "POST")

	var createData createPayment
	err := c.BindJSON(&createData)
	if err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, "Create Parameters not specified")
			return
		}
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// if there is a date, convert it
	if createData.Date == "" {
		c.JSON(http.StatusBadRequest, "Date not specified")
		return
	}
	ttime, err := time.Parse("2006-01-02", createData.Date)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if createData.Value == "" {
		c.JSON(http.StatusBadRequest, "Value not specified")
		return
	}
	value, err := strconv.ParseFloat(createData.Value, 64)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if createData.ExpenseID <= 0 {
		c.JSON(http.StatusBadRequest, "Specified expense is invalid.")
		return
	}
	expID := createData.ExpenseID

	exp, err := r.dm.CreatePayment(ttime, value, expID)
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, exp)
}
