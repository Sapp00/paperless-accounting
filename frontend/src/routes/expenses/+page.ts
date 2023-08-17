
import { onMount } from "svelte";

type Expense = {
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

type Payment = {
    ID: number,
    Date: string,
    Value: number,
    ExpenseID: number
}

type ChartEntry = {
    date: string,
    category: string,
    value: number
};

/** @type {import('./$types').PageLoad} */
export async function load({ fetch }) {
    const urlParams = new URLSearchParams(window.location.search);
    const fromTime = urlParams.get('from');
    const toTime = urlParams.get('to');

    console.log(`from: ${fromTime} to ${toTime}`);

    try{
        const res = await fetch(`http://localhost:8080/expenses?` + new URLSearchParams({
            from: fromTime!,
            to: toTime!,
        }));
        const expenses: Expense[] = await res.json();

        const resP = await fetch(`http://localhost:8080/payments?` + new URLSearchParams({
            from: fromTime!,
            to: toTime!,
        }));
        const payments: Payment[] = await resP.json();

        const itemsE = new Map<string, ChartEntry>();
        const itemsP = new Map<string, ChartEntry>();
        let j=0;


        // process payments data for chart
        payments.forEach(p => {
            let v = itemsP.get(p.Date)
            if (v == undefined){
                let vn = {
                    date: p.Date,
                    category: "payment",
                    value: p.Value
                };
                itemsP.set(p.Date, vn);
            } else {
                v!.value += p.Value
            }
        });

        payments.forEach(p => {
            let v = itemsP.get(p.Date)
            if (v == undefined){
                let vn = {
                    date: p.Date,
                    category: "payment",
                    value: p.Value
                };
                itemsP.set(p.Date, vn);
            } else {
                v!.value += p.Value
            }
        });

        expenses.forEach(e => {
            let v = itemsE.get(e.Date)
            if (v == undefined){
                let vn = {
                    date: e.Date,
                    category: "expense",
                    value: e.Value
                };
                itemsE.set(e.Date, vn);
            } else {
                v!.value += e.Value
            }
        });

        // combine both maps to one array of values
        const items = Array.from(itemsE.values()).concat( Array.from(itemsP.values()) );
        
        console.log(items);

        return {
            items: items,
            expenses: expenses,
            error: null,
            fromDate: new Date(fromTime!),
            toDate: new Date(toTime!),
        };
    }
    catch(e){
        console.log("error: "+e);
        return {
            expenses: [],
            error: "Cannot establish connection",
            fromDate: null,
            toDate: null,
        }
    }
}