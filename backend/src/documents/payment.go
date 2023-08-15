package documents

import (
	"context"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/paperless"
	"time"
)

type Payment struct {
	// from database
	ID        int
	Date      paperless.PaperlessTime
	Value     float32
	ExpenseID int
}

func (m *DocumentMgr) GetPayment(id int) (*Payment, error) {
	ctx := context.Background()
	paymentDB, err := m.db.GetPayment(ctx, int64(id))

	if err != nil {
		return nil, err
	}

	payment := &Payment{
		ID:        id,
		Date:      *paperless.NewPaperlessTime(paymentDB.Paiddate),
		Value:     float32(paymentDB.Price),
		ExpenseID: int(paymentDB.Expenseid),
	}

	return payment, nil
}

func (m *DocumentMgr) GetPaymentsBetween(from, to time.Time) ([]*Payment, error) {
	ctx := context.Background()

	paymentDB, err := m.db.GetPaymentsBetween(ctx, database.GetPaymentsBetweenParams{
		Paiddate:   from,
		Paiddate_2: to,
	})

	if err != nil {
		return nil, err
	}

	payments := make([]*Payment, len(paymentDB))

	for i, p := range paymentDB {
		payment := &Payment{
			ID:        int(p.ID),
			Date:      *paperless.NewPaperlessTime(p.Paiddate),
			Value:     float32(p.Price),
			ExpenseID: int(p.Expenseid),
		}

		payments[i] = payment
	}

	return payments, nil
}

func (m *DocumentMgr) GetPaymentsByExpenseID(id int) ([]*Payment, error) {
	ctx := context.Background()

	paymentDB, err := m.db.GetPaymentsByExpense(ctx, int64(id))

	if err != nil {
		return nil, err
	}

	payments := make([]*Payment, len(paymentDB))

	for i, p := range paymentDB {
		payment := &Payment{
			ID:        int(p.ID),
			Date:      *paperless.NewPaperlessTime(p.Paiddate),
			Value:     float32(p.Price),
			ExpenseID: int(p.Expenseid),
		}

		payments[i] = payment
	}

	return payments, nil
}
