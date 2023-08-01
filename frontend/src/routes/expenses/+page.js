
import { onMount } from "svelte";

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {
    try{
        const res = await fetch("http://localhost:8080/expenses");
        const item = await res.json();

        return {
            items: item
        };
    }
    catch(e){
        return {
            items: [],
            error: "Cannot establish connection"
        }
    }
}