package gql_resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"log"
	"sapp/paperless-accounting/documents"
	"sapp/paperless-accounting/graph/gql_generated"
	"sapp/paperless-accounting/graph/gql_types"
	"sort"
)

// PaperlessInt is the resolver for the paperlessInt field.
func (r *incomeResolver) PaperlessInt(ctx context.Context, obj *documents.Income) (int, error) {
	panic(fmt.Errorf("not implemented: PaperlessInt - paperlessInt"))
}

// CreatePayment is the resolver for the createPayment field.
func (r *mutationResolver) CreatePayment(ctx context.Context, input gql_types.NewPayment) (*documents.Payment, error) {
	panic(fmt.Errorf("not implemented: CreatePayment - createPayment"))
}

// UpdatePayment is the resolver for the updatePayment field.
func (r *mutationResolver) UpdatePayment(ctx context.Context, input gql_types.UpdatePayment) (*documents.Payment, error) {
	panic(fmt.Errorf("not implemented: UpdatePayment - updatePayment"))
}

// DeletePayment is the resolver for the deletePayment field.
func (r *mutationResolver) DeletePayment(ctx context.Context, id int) (*documents.Payment, error) {
	panic(fmt.Errorf("not implemented: DeletePayment - deletePayment"))
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, input gql_types.UpdateExpense) (*documents.Expense, error) {
	panic(fmt.Errorf("not implemented: UpdateExpense - updateExpense"))
}

// UpdateIncome is the resolver for the updateIncome field.
func (r *mutationResolver) UpdateIncome(ctx context.Context, input gql_types.UpdateIncome) (*documents.Income, error) {
	panic(fmt.Errorf("not implemented: UpdateIncome - updateIncome"))
}

// Value is the resolver for the value field.
func (r *paymentResolver) Value(ctx context.Context, obj *documents.Payment) (float64, error) {
	panic(fmt.Errorf("not implemented: Value - value"))
}

// GetIncomes is the resolver for the getIncomes field.
func (r *queryResolver) GetIncomes(ctx context.Context) ([]*documents.Income, error) {
	panic(fmt.Errorf("not implemented: GetIncomes - getIncomes"))
}

// GetExpenses is the resolver for the getExpenses field.
func (r *queryResolver) GetExpenses(ctx context.Context) ([]*documents.Expense, error) {
	return r.Dm.GetExpenses()
}

// GetPayments is the resolver for the getPayments field.
func (r *queryResolver) GetPayments(ctx context.Context) ([]*documents.Payment, error) {
	panic(fmt.Errorf("not implemented: GetPayments - getPayments"))
}

// GetExpensesBetween is the resolver for the getExpensesBetween field.
func (r *queryResolver) GetExpensesBetween(ctx context.Context, from string, to string) ([]*gql_types.ChartEntry, error) {
	all_expenses, err := r.Dm.GetExpensesBetween(from, to)
	if err != nil {
		log.Fatalf("Error occured when retrieving expenses: %s\n", err.Error())
	}

	// TODO: implement this
	all_payments := all_expenses

	fmt.Printf("implement me! all_payments\n")

	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
		return nil, err
	}

	sort.SliceStable(all_expenses, func(i, j int) bool {
		return all_expenses[i].Created_date.Before(all_expenses[j].Created_date.Time)
	})

	/*
		sort.SliceStable(all_payments, func(i, j int) bool {
			return all_payments[i].Created_date.Before(all_payments[j].Created_date.Time)
		})*/

	out := make([]*gql_types.ChartEntry, len(all_expenses)+len(all_payments))

	// create entries for expenses
	var expense_sum float64 = 0
	i := 0
	for _, e := range all_expenses {
		// TODO: change created date! needs to be based on paid_date which is retrieved from the database
		e_exp_val := float64(int(e.Title[0]) * 20)
		e_exp_date := e.Created_date

		// update expense
		expense_sum += e_exp_val
		if i != 0 && out[i-1].Date == e_exp_date {
			out[i-1].Value = expense_sum
		} else {
			out[i] = &gql_types.ChartEntry{Date: e_exp_date, Category: "expense", Value: expense_sum}
			i++
		}
	}
	// is empty or has an expense?
	if len(all_expenses) > 0 && out[i] == nil {
		i++
	}
	// create entries for payments
	var paid_sum float64 = 0
	for _, e := range all_payments {
		// TODO: change created date! needs to be based on paid_date which is retrieved from the database
		var e_paid_val float64 = float64(e.Value)
		e_paid_date := e.Date

		// update expense
		paid_sum += e_paid_val
		if i != 0 && out[i-1].Date == e_paid_date {
			out[i-1].Value = paid_sum
		} else {
			out[i] = &gql_types.ChartEntry{Date: e_paid_date, Category: "payment", Value: paid_sum}
			i++
		}
	}
	// is empty or has an expense?
	if len(all_payments) > 0 && out[i] != nil {
		i--
	}

	outputTruncated := out[:i]

	return outputTruncated, nil
}

// GetIncome is the resolver for the getIncome field.
func (r *queryResolver) GetIncome(ctx context.Context, id int) (*documents.Income, error) {
	panic(fmt.Errorf("not implemented: GetIncome - getIncome"))
}

// GetExpense is the resolver for the getExpense field.
func (r *queryResolver) GetExpense(ctx context.Context, id int) (*documents.Expense, error) {
	panic(fmt.Errorf("not implemented: GetExpense - getExpense"))
}

// GetPayment is the resolver for the getPayment field.
func (r *queryResolver) GetPayment(ctx context.Context, id int) (*documents.Payment, error) {
	panic(fmt.Errorf("not implemented: GetPayment - getPayment"))
}

// Income returns gql_generated.IncomeResolver implementation.
func (r *Resolver) Income() gql_generated.IncomeResolver { return &incomeResolver{r} }

// Mutation returns gql_generated.MutationResolver implementation.
func (r *Resolver) Mutation() gql_generated.MutationResolver { return &mutationResolver{r} }

// Payment returns gql_generated.PaymentResolver implementation.
func (r *Resolver) Payment() gql_generated.PaymentResolver { return &paymentResolver{r} }

// Query returns gql_generated.QueryResolver implementation.
func (r *Resolver) Query() gql_generated.QueryResolver { return &queryResolver{r} }

type incomeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type paymentResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
