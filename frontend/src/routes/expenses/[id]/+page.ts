
import { onMount } from "svelte";

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

type Payment = {
    ID: number,
    Date: string,
    Value: number,
    ExpenseID: number
}

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {

    try{
        const res = await fetch(`http://localhost:8080/expense/${params.id}`);
        const expense: Expense[] = await res.json();

        const resP = await fetch(`http://localhost:8080/payments?` + new URLSearchParams({
            expense: params.id
        }));
        const payments: Payment[] = await resP.json();


        return {
            expense: expense
        };
    }
    catch(e){
        console.log("error: "+e);
        return {
            expense: null,
            error: "Cannot establish connection",
            fromDate: null,
            toDate: null,
        }
    }
}