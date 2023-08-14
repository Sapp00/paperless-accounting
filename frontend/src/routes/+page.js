
import { onMount } from "svelte";

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {
    try{
        const res = await fetch("http://localhost:8080/expensesTODO/");
        const item = await res.json();


        return {
            pokemons: item
        };
    }
    catch(e){
        return {
            pokemons: [],
            error: "Cannot establish connection"
        }
    }
}