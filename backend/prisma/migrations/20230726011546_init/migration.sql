-- CreateTable
CREATE TABLE "Expense" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "price" DECIMAL NOT NULL,
    "paid" BOOLEAN NOT NULL,
    "paidDate" DATETIME
);
