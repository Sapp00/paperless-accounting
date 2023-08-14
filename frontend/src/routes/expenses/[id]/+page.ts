import { onMount } from "svelte";

export type ItemType = {
    Date: string,
    Value: number,
    PaperlessID: number,
    Correspondent: number,
    Title: string,
    Content: string,
    Tags: number[],
    Created_date: string,
};

export interface DataType {
    item: ItemType,
    error: string,
}

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {
    try{
        if (params.id != ''){
            const res = await fetch(`http://localhost:8080/expenses/${params.year}/${params.id}`);
            const item: ItemType = await res.json();
    
            return {
                item: item,
                error: ""
            } as DataType;
        } else {
            return {
                item: null as unknown as ItemType,
                error: "No year parameter defined."
            } as DataType
        }
    }
    catch(e){
        return {
            item: null as unknown as ItemType,
            error: "Cannot establish connection"
        } as DataType
    }
}