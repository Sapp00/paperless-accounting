
import { onMount } from "svelte";
import type{ Expense, Payment, ChartEntry} from "$lib/types";

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

        if (payments != null){
            console.log(`payments:${payments.length}`);
            for (let i=0; i < payments.length; ){
                let p = payments[i];
                let e = expenses[j];

                if (p.ExpenseID == e.PaperlessID){
                    // exists? add
                    if(e.paidValue != null){
                        e.paidValue += p.Value;
                    } else {
                        e.paidValue = p.Value;
                    }
                    // add payment to chart and aggregate
                    if ( itemsP.has(p.Date) ){
                        itemsP.get(p.Date)!.value += p.Value;
                    } else {
                        itemsP.set(p.Date, {
                            date: p.Date,
                            category: "payment",
                            value: p.Value
                        });
                    }
                    ++i;
                // next
                } else {
                    // add expense to chart and aggregate
                    if ( itemsE.has(e.Date) ){
                        itemsE.get(e.Date)!.value += e.Value;
                    } else {
                        itemsE.set(e.Date, {
                            date: e.Date,
                            category: "expense",
                            value: e.Value
                        });
                    }
                    j++;
                }
            }
        }

        console.log(`j:${j}`);

        // add rest of the expenses
        for(; j < expenses.length; ++j){
            let e = expenses[j];
            if ( itemsE.has(e.Date) ){
                itemsE.get(e.Date)!.value += e.Value;
            } else {
                itemsE.set(e.Date, {
                    date: e.Date,
                    category: "expense",
                    value: e.Value
                });
            }
        }

        // combine both maps to one array of values
        let items = Array.from(itemsE.values()).concat( Array.from(itemsP.values()) );

        items.sort(function(a,b: ChartEntry) {return new Date(a.date).valueOf() - new Date(b.date).valueOf()});
        
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