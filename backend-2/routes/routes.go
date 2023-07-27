package routes

import (
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/routes/expenses"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	conf   *config.Config
	router *gin.Engine
}

func New(conf *config.Config) (*Routes, error) {
	r := Routes{
		conf: conf,
	}

	r.router = gin.Default()

	return &r, nil
}

func (r Routes) Setup() error {
	exp := expenses.NewRouter(r.conf)
	r.router.GET("/expenses", exp.GetAllExpenses)

	err := r.router.Run("localhost:8080")

	return err
}
