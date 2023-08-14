
import { onMount } from "svelte";

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
        const item = await res.json();

        return {
            items: item,
            error: null,
            fromDate: new Date(fromTime!),
            toDate: new Date(toTime!),
        };
    }
    catch(e){
        return {
            items: [],
            error: "Cannot establish connection",
            fromDate: null,
            toDate: null,
        }
    }
}