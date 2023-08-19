<script lang="ts">
    import type {Payment, Expense} from "$lib/types";

    import { onMount } from "svelte";

    export let payment: Payment | null;
    export let expense: Expense | null;
    var expenses: Expense[];

   // export let val;
    
    onMount(async () => {
        if(payment != null){
            const res = await fetch(`http://localhost:8080/expense/${payment.ExpenseID}`);
            expense = await res.json();  
        } else if(expense == null){
            const res = await fetch("http://localhost:8080/expenses");
            expenses = await res.json();
        }
    });

    let dateInput;
    let priceInput;
    function SendForm(){
        
    }
</script>

<div class={$$restProps.class}>
    {#if $$restProps.head}
        <div class="flex justify-center">
            <h1 class="h3">
                {$$restProps.head}
            </h1>
        </div>
    {/if}
    <label class="label m-4">
        <span>Date</span>
        <input class="input" type="date" bind:this={dateInput} 
            value={expense ? expense.Date : null}
        />
    </label>
    <label class="label m-4">
        <span>Price</span>
        <input class="input" type="number" min="0" step="0.01" bind:this={priceInput}
            value={expense ? expense.Value : null}
        />
    </label>
    {#if !expense}
    <label class="label m-4">
        <span>Expense</span>
        <select class="select" value={expense}>
            {#each expenses as e}
                <option value={e.PaperlessID}>{e.Title}</option>
            {/each}
        </select>
    </label>
    {/if}
    <div class="flex justify-center items-center m-4">
        <button type="button" class="btn variant-filled w-full" on:click={SendForm}>Update</button>
    </div>
</div>