package expenses

import (
	"net/http"
	"sapp/paperless-accounting/config"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []Expense{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (r *ExpenseRouter) GetAllExpenses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func NewRouter(conf *config.Config) *ExpenseRouter {
	r := ExpenseRouter{
		conf: conf,
	}

	return &r
}
