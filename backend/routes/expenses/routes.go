package expenses

import (
	"log"
	"net/http"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/paperless"

	"github.com/gin-gonic/gin"
)

func (r *ExpenseRouter) GetAllExpenses(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, albums)

	all_expenses, err := r.paperless.PaperlessDocumentQuery("tag:" + r.conf.PAPERLESS_EXPENSE_TAG)

	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	// TODO: here is potential for performance optimization
	expense_chart_map := make(map[paperless.PaperlessTime]chartEntry)
	paid_chart_map := make(map[paperless.PaperlessTime]chartEntry)
	var expense_sum float32 = 0
	var paid_sum float32 = 0

	// create entries
	for _, e := range all_expenses {
		// TODO: change created date! needs to be based on paid_date which is retrieved from the database
		e_exp_val := float32(int(e.Title[0]) * 20)
		var e_paid_val float32 = 0
		if int(e_paid_val/3) != 0 {
			e_paid_val = e_exp_val
		}
		e_exp_date := e.Created_date
		e_paid_date := e_exp_date

		// update expense
		expense_sum += e_exp_val
		if val, ok := expense_chart_map[e_exp_date]; ok {
			val.Value = expense_sum
		} else {
			expense_chart_map[e_exp_date] = chartEntry{Date: e_exp_date, Category: "expense", Value: expense_sum}
		}

		// update payment
		paid_sum += e_paid_val
		if val, ok := paid_chart_map[e_paid_date]; ok {
			val.Value = paid_sum
		} else {
			paid_chart_map[e_paid_date] = chartEntry{Date: e_paid_date, Category: "payment", Value: paid_sum}
		}
	}

	// merge maps
	out := make([]chartEntry, len(expense_chart_map)+len(paid_chart_map))
	i := 0
	for _, e := range expense_chart_map {
		out[i] = e
		i++
	}
	for _, e := range paid_chart_map {
		out[i] = e
		i++
	}

	c.JSON(http.StatusOK, out)

}

func NewRouter(conf *config.Config) (*ExpenseRouter, error) {
	p, err := paperless.Init(conf)
	if err != nil {
		return nil, err
	}

	r := ExpenseRouter{
		conf:      conf,
		paperless: p,
	}

	return &r, nil
}
