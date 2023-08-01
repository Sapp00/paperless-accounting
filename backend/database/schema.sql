CREATE TABLE expenses (
  id   INTEGER PRIMARY KEY,
  price DECIMAL(12,2),
  expenseDate date
);

CREATE TABLE payments (
  id   INTEGER PRIMARY KEY,
  expenseID INTEGER NOT NULL,
  price DECIMAL(12,2)    NOT NULL,
  paidDate date,
  FOREIGN KEY(expenseID) REFERENCES expenses(id)
);