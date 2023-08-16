package documents

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sapp/paperless-accounting/database"
	"sapp/paperless-accounting/paperless"
	"time"

	"github.com/redis/go-redis/v9"
)

type Expense struct {
	// from database
	Date  paperless.PaperlessTime
	Value float64

	// from paperless
	PaperlessID   int
	Correspondent int
	Title         string
	Content       string
	Tags          []int
	Created_date  paperless.PaperlessTime
}

func (m *DocumentMgr) GetExpense(id int) (*Expense, error) {
	ctx := context.Background()

	val, err := m.client.ZRangeByScore(ctx, "expenses", &redis.ZRangeBy{
		Min:   fmt.Sprint(id),
		Max:   fmt.Sprint(id),
		Count: 1,
	}).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	} else if len(val) == 0 {
		return nil, nil
	}

	var res Expense
	err = json.Unmarshal([]byte(val[0]), &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *DocumentMgr) GetExpensesBetween(fromTimeStr string, toTimeStr string) ([]*Expense, error) {
	ctx := context.Background()

	var err error
	var fromTime time.Time
	if fromTimeStr == "-1" {
		fromTime = time.Unix(0, 0)
	} else {
		fromTime, err = time.Parse("2006-01-02", fromTimeStr)
		if err != nil {
			return nil, err
		}
	}
	var toTime time.Time
	if toTimeStr == "0" {
		toTime = time.Now()
	} else {
		toTime, err = time.Parse("2006-01-02", toTimeStr)
		if err != nil {
			return nil, err
		}
	}

	val, err := m.client.ZRangeByScore(ctx, "expenses_by_date", &redis.ZRangeBy{
		Min: fmt.Sprint(fromTime.Unix()),
		Max: fmt.Sprint(toTime.Unix()),
	}).Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}

	log.Printf("From %v to %v: Got %d results\n", fromTime.Unix(), toTime.Unix(), len(val))

	var res []*Expense

	for _, v := range val {
		var p []Expense
		err = json.Unmarshal([]byte(v), &p)

		if err != nil {
			return nil, err
		}
		for _, r := range p {
			res = append(res, &r)
		}
	}

	return res, nil
}

func (m *DocumentMgr) GetExpenses() ([]*Expense, error) {
	return m.GetExpensesBetween("-1", "0")
}

func (m *DocumentMgr) UpdateExpense(id int, date *paperless.PaperlessTime, value *float64) (*Expense, error) {
	exp, err := m.GetExpense(id)
	if err != nil {
		return nil, errors.New("expense cannot be updated because it does not exist. create it first")
	}
	if date != nil {
		exp.Date = *date
	}
	if value != nil {
		exp.Value = *value
	}

	fmt.Printf("changing date of %d to %v and value to %v\n", id, *date, *value)

	ctx := context.Background()
	m.db.UpdateExpense(ctx, database.UpdateExpenseParams{Price: sql.NullFloat64{Valid: true, Float64: exp.Value}, Expensedate: exp.Date.Time, ID: int64(id)})

	ej, err := json.Marshal(&exp)
	if err != nil {
		return nil, err
	}

	// update expenses
	err = m.client.ZRem(ctx, "expenses", redis.Z{Score: float64(id), Member: ej}).Err()
	if err != nil {
		return nil, err
	}
	err = m.client.ZAdd(ctx, "expenses", redis.Z{Score: float64(id), Member: ej}).Err()
	if err != nil {
		return nil, err
	}

	// update by date
	err = m.client.ZRem(ctx, "expenses_by_date", redis.Z{Score: float64(exp.Date.Unix()), Member: ej}).Err()
	if err != nil {
		return nil, err
	}
	err = m.client.ZAdd(ctx, "expenses_by_date", redis.Z{Score: float64(exp.Date.Unix()), Member: ej}).Err()
	if err != nil {
		return nil, err
	}

	return exp, nil
}
