export type Expense = {
    Date: string,
    Value: number,
    PaperlessID: number,
    Correspondent: number,
    Title: string,
    Content: string,
    Tags: number[],
    Created_date: string,
    paidValue?: number
}

export type Payment = {
    ID: number,
    Date: string,
    Value: number,
    ExpenseID: number
}

export type ChartEntry = {
    date: string,
    category: string,
    value: number
};