package expenses

import (
	"fmt"
	"log"
	"net/http"
	"sapp/paperless-accounting/config"
	"sapp/paperless-accounting/paperless"
	"sort"

	"github.com/gin-gonic/gin"
)

func (r *ExpenseRouter) GetAllExpenses(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, albums)

	all_expenses, err := r.paperless.GetDocuments(paperless.Expense, "2023")
	if err != nil {
		log.Fatalf("Error occured when retrieving expenses: %s\n", err.Error())
	}
	// TODO: implement this
	all_payments := all_expenses

	fmt.Printf("implement me! all_payments\n")

	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	sort.SliceStable(all_expenses, func(i, j int) bool {
		return all_expenses[i].Created_date.Before(all_expenses[j].Created_date.Time)
	})

	/*
		sort.SliceStable(all_payments, func(i, j int) bool {
			return all_payments[i].Created_date.Before(all_payments[j].Created_date.Time)
		})*/

	out := make([]chartEntry, len(all_expenses)+len(all_payments))

	// create entries for expenses
	var expense_sum float32 = 0
	i := 0
	for _, e := range all_expenses {
		// TODO: change created date! needs to be based on paid_date which is retrieved from the database
		e_exp_val := float32(int(e.Title[0]) * 20)
		e_exp_date := e.Created_date

		// update expense
		expense_sum += e_exp_val
		if i != 0 && out[i-1].Date == e_exp_date {
			out[i-1].Value = expense_sum
		} else {
			out[i] = chartEntry{Date: e_exp_date, Category: "expense", Value: expense_sum}
			i++
		}
	}
	// is empty or has an expense?
	if out[i] == (chartEntry{}) {
		i++
	}
	// create entries for payments
	var paid_sum float32 = 0
	for _, e := range all_expenses {
		// TODO: change created date! needs to be based on paid_date which is retrieved from the database
		e_exp_val := float32(int(e.Title[0]) * 20)
		var e_paid_val float32 = 0
		if (int(e_exp_val) % 3) != 0 {
			e_paid_val = e_exp_val
		}
		e_paid_date := e.Created_date

		// update expense
		paid_sum += e_paid_val
		if i != 0 && out[i-1].Date == e_paid_date {
			out[i-1].Value = paid_sum
		} else {
			out[i] = chartEntry{Date: e_paid_date, Category: "payment", Value: paid_sum}
			i++
		}
	}
	// is empty or has an expense?
	if out[i] != (chartEntry{}) {
		i--
	}

	outputTruncated := out[:i]

	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.JSON(http.StatusOK, outputTruncated)

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
