<script lang="ts">
    import type {Expense} from "$lib/types";

    import { onMount } from "svelte";

    export let expense: Expense;
    var expenses: Expense[];
    
    onMount(async () => {
        if(expenseID == null){
            const res = await fetch(`http://localhost:8080/expenses?` + new URLSearchParams({
                from: fromTime!,
                to: toTime!,
            }));
            expenses = await res.json();
        }
    });

    let payment;
    let dayInput;
    let priceInput;
</script>

<label class="label m-4">
    <span>Date</span>
    <input class="input" type="date" value={expense.Date} bind:this={dateInput}/>
</label>
<label class="label m-4">
    <span>Price</span>
    <input class="input" type="number" min="0" step="0.01" value={expense.Value} bind:this={priceInput}/>
</label>
<label class="label m-4">
    <span>Expense</span>
    {#if !expense}
    <select class="select" value={expense}>
        {#each expenses as e}
            <option value={correspondent.id}>{correspondent.name}</option>
        {/each}
    </select>
</label>
<div class="flex justify-center items-center m-4">
    <button type="button" class="btn-lg variant-filled" on:click={SendForm}>Update</button>
</div>