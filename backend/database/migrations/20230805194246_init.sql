-- +goose Up
-- +goose StatementBegin
CREATE TABLE expenses (
  id   INTEGER PRIMARY KEY,
  price DECIMAL(12,2),
  expenseDate date NOT NULL
);

CREATE TABLE payments (
  id   INTEGER PRIMARY KEY,
  expenseID INTEGER NOT NULL,
  price DECIMAL(12,2)    NOT NULL,
  paidDate date NOT NULL,
  FOREIGN KEY(expenseID) REFERENCES expenses(id)
);

CREATE TABLE incomes (
  id   INTEGER PRIMARY KEY,
  price DECIMAL(12,2),
  incomeDate date NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE expenses;
DROP TABLE payments;
DROP TABLE incomes;
-- +goose StatementEnd
