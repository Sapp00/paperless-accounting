package routes

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/routes/expenses"
	"sapp/paperless-accounting/routes/payments"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	conf   *config.Config
	router *gin.Engine
	dm     *documents.DocumentMgr
}

func New(conf *config.Config) (*Routes, error) {
	r := Routes{
		conf: conf,
	}

	if !conf.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r.router = gin.Default()

	dm, err := documents.NewManager(conf)
	if err != nil {
		return nil, err
	}
	r.dm = dm

	return &r, nil
}

func (r Routes) Setup() error {
	exp, err := expenses.New(r.conf, r.dm)
	if err != nil {
		return err
	}

	r.router.GET("/expenses", exp.GetExpensesBetween)
	r.router.GET("/expenses/:id/payments", exp.GetPayments)
	r.router.GET("/expenses/:id", exp.GetExpense)

	pay, err := payments.New(r.conf, r.dm)
	if err != nil {
		return err
	}
	r.router.GET("/payments", pay.GetPaymentsBetween)
	r.router.GET("/payments/:id", pay.GetPayment)

	err = r.router.Run("localhost:8080")
	return err
}
