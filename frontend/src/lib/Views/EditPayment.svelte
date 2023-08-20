<script lang="ts">
    import { onMount } from "svelte";
    
    import type {Payment, Expense} from "$lib/types";
    import {isValidFloat2D} from '$lib/helper/validation';

    import type { ToastSettings } from '@skeletonlabs/skeleton';
    import { modalStore, toastStore } from '@skeletonlabs/skeleton';

    export let payment: Payment | null;
    export let expense: Expense | null;
    var expenses: Expense[];

   // export let val;
    
    onMount(async () => {
        if(payment != null){
            const res = await fetch(`http://localhost:8080/expenses/${payment.ExpenseID}`);
            expense = await res.json();  
        } else if(expense == null){
            const res = await fetch("http://localhost:8080/expenses");
            expenses = await res.json();
        }
    });

    let actionPhrase = payment ? "updated" : "added";
    let buttonPhrase = payment ? "Update" : "Create";
    const tErr: ToastSettings = {
        message: 'Invalid Input, please correct the highlighted fields.'
    }
    const tSucc: ToastSettings = {
        message: `Successfully ${actionPhrase} the payment.`
    }

    let dateInput: HTMLInputElement;
    let priceInput: HTMLInputElement;
    let expenseInput: HTMLSelectElement;
    function SendForm(){
        if (isValidFloat2D(priceInput.value)){
            priceInput.classList.remove("input-error");

            let eID: number = expense ? expense.PaperlessID : parseInt(expenseInput.value);

            var req: Promise<Response>;
            if(payment == null){
                let body = { date: dateInput.value, value: priceInput.value, expense: 0};
                if(expense == null) {
                    body.expense = parseInt( expenseInput.value );
                } else {
                    body.expense = expense.PaperlessID;
                }
                let opt = {
                    method: 'POST',
                    body: JSON.stringify( body )
                };
                req = fetch(`http://localhost:8080/payments`, opt);
            } else {
                let opt = {
                    method: 'POST',
                    body: JSON.stringify( { date: dateInput.value, value: priceInput.value} )
                };
                req = fetch(`http://localhost:8080/payments/${payment!.ID}/`, opt);
            }


            req
            .then( res => res.json())
            .then( res => {
                if( res != null){
                    console.log(res);
                    var paymentOutput: Payment;
                    if(payment){
                        paymentOutput = {ID: payment.ID, Date: dateInput.value, ExpenseID: eID, Value: parseFloat(priceInput.value)};
                    } else {
                        paymentOutput = {ID: res.ID, Date: dateInput.value, ExpenseID: eID, Value: parseFloat(priceInput.value)};
                    }
                    $modalStore[0].response!(paymentOutput);
                    toastStore.trigger(tSucc);
                }
            })
            .catch(error => {
                console.error(error);
            });
        } else {
            priceInput.classList.add("input-error");
            toastStore.trigger(tErr);
        }
    }

    function DeleteForm(){
        fetch(`http://localhost:8080/payments/${payment!.ID}`, {
            method: 'DELETE'
        })
        .then( res => res.json())
        .then( res => {
            const tSuccMsg: ToastSettings = {
                message: `Successfully deleted the payment.`
            };
            // todo change response type -> need to detect deletion. send "action: delete|add|null, object: Payment"
            $modalStore[0].response!(null);
        } )
        .catch( error => {
            const tErrMsg: ToastSettings = {
                message: `Error! Could not delete the payment: ${error}`
            };
            toastStore.trigger(tErrMsg);
        })
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
            value={payment ? payment.Date : (expense ? expense.Date : null)}
        />
    </label>
    <label class="label m-4">
        <span>Price</span>
        <input class="input" type="number" min="0" step="0.01" bind:this={priceInput}
            value={payment ? payment.Value : (expense ? expense.Value : null)}
        />
    </label>
    {#if !expense}
    <label class="label m-4">
        <span>Expense</span>
        <select class="select" bind:this={expenseInput}>
            {#each expenses as e}
                <option value={e.PaperlessID}>{e.Title}</option>
            {/each}
        </select>
    </label>
    {/if}
    <div class="items-center m-4">
        <button type="button" class="btn variant-filled-primary w-full my-4" on:click={SendForm}>{buttonPhrase}</button>
        {#if payment}
            <button type="button" class="btn variant-ghost-error w-full" on:click={DeleteForm}>Delete</button>
        {/if}
    </div>
</div>