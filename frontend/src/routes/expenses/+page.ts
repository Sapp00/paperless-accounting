
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


        let items: ChartEntry[] = [];
        const itemsE = new Map<number, Expense>();

        expenses.sort(function(a,b){return Date.parse(a.Date) - Date.parse(b.Date)});
        payments.sort(function(a,b){return Date.parse(a.Date) - Date.parse(b.Date)});

        let esum = 0;
        let eprev: Expense|null = null;
        expenses.forEach(e => {
            if (eprev == null){
                eprev = e;
            } else if (eprev.Date != e.Date){
                items.push({date: eprev.Date, category: "expense", value: esum });
                eprev = e;
            }
            esum += e.Value;
            
            e.paidValue = 0;
            itemsE.set(e.PaperlessID, e);
        });
        
        let psum = 0;
        let pprev: Payment|null = null;
        payments.forEach(p => {
            if(pprev == null){
                pprev = p;
            } else if(pprev.Date != p.Date){
                items.push({date: pprev.Date, category: "payment", value: psum});
                pprev = p;
            }
            psum += p.Value;
            itemsE.get(p.ExpenseID)!.paidValue! += p.Value;
        });

        // combine both maps to one array of values
        //let items = Array.from(itemsE.values()).concat( Array.from(itemsP.values()) );

       // items.sort(function(a,b: ChartEntry) {return new Date(a.date).valueOf() - new Date(b.date).valueOf()});
        

        let expensesOut = Array.from(itemsE.values());
        console.log(expensesOut);

        return {
            items: items,
            expenses: expensesOut,
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