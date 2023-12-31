package routes

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/routes/correspondents"
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

	r.router.Use(corsMiddleware())

	dm, err := documents.NewManager(conf)
	if err != nil {
		return nil, err
	}
	r.dm = dm

	return &r, nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (r Routes) Setup() error {
	exp, err := expenses.New(r.conf, r.dm)
	if err != nil {
		return err
	}

	r.router.GET("/expenses", exp.GetExpensesBetween)
	r.router.GET("/expenses/:id/payments", exp.GetPayments)
	r.router.GET("/expenses/:id", exp.GetExpense)
	r.router.GET("/expenses/:id/thumb/", exp.GetExpenseThumb)

	r.router.POST("/expenses/:id/", exp.UpdateExpense)

	pay, err := payments.New(r.conf, r.dm)
	if err != nil {
		return err
	}
	r.router.GET("/payments", pay.GetPaymentsBetween)
	r.router.GET("/payments/:id", pay.GetPayment)

	r.router.POST("/payments", pay.CreatePayment)
	r.router.POST("/payments/:id", pay.UpdatePayment)

	r.router.DELETE("/payments/:id", pay.DeletePayment)

	corr, err := correspondents.New(r.conf, r.dm)
	if err != nil {
		return err
	}
	r.router.GET("/correspondents", corr.GetCorrespondents)
	r.router.GET("/correspondents/:id", corr.GetCorrespondent)

	err = r.router.Run("localhost:8080")
	return err
}
