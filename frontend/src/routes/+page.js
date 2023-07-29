
import { onMount } from "svelte";

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {
    const res = await fetch("https://pokeapi.co/api/v2/pokemon?limit=151");
    const item = await res.json();

    return {
        pokemons: item['results']
    };
}