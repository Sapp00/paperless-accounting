
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

export type Payment = {
    ID: number,
    Date: string,
    Value: number,
    ExpenseID: number
}

export type Correspondent = {
    id: number,
    name: string
}

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {

    try{
        const res = await fetch(`http://localhost:8080/expenses/${params.id}`);
        const expense: Expense = await res.json();

        const resP = await fetch(`http://localhost:8080/expenses/${params.id}/payments`);
        const payments: Payment[] = await resP.json();

        const resCorr = await fetch(`http://localhost:8080/correspondents`);
        const correspondents: Correspondent[] = await resCorr.json();


        return {
            expense: expense,
            payments: payments,
            correspondents: correspondents
        };
    }
    catch(e){
        console.log("error: "+e);
        return {
            expense: null,
            payments: [] as Payment[],
            correspondents: [] as Correspondent[],
            error: "Cannot establish connection",
            fromDate: null,
            toDate: null,
        }
    }
}