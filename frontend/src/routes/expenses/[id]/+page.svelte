<script lang="ts">
    import type { PageData } from './$types';  
    import { onMount } from "svelte";
	import type { Expense, Payment } from './+page';
    import {isValidFloat2D} from '$lib/helper/validation';

    import { Table, tableMapperValues, type TableSource, toastStore, type ToastSettings, TabGroup, Tab } from '@skeletonlabs/skeleton';

	export let data: PageData;
    const { expense, payments, correspondents, error } = data;

    import MdEdit from 'svelte-icons/md/MdEdit.svelte'
    import MdPayment from 'svelte-icons/md/MdPayment.svelte';


    let tabSet: number = 0;
    // Update
    let priceInput: HTMLInputElement;
    let dateInput: HTMLInputElement;
    const tErr: ToastSettings = {
        message: 'Invalid Input, please correct the highlighted fields.'
    }
    const tSucc: ToastSettings = {
        message: 'Successfully updated the expense.'
    }
    function SendForm(){
        if (isValidFloat2D(priceInput.value)){
            priceInput.classList.remove("input-error");

            let opt = {
                method: 'POST',
                body: JSON.stringify( { date: dateInput.value, value: priceInput.value} )
            }
            fetch(`http://localhost:8080/expenses/${expense!.PaperlessID}/`, opt)
            .then( res => res.json())
            .then( res => {
                if( res != null){
                    toastStore.trigger(tSucc);
                }
            });
        } else {
            priceInput.classList.add("input-error");
            toastStore.trigger(tErr);
        }
    }

    // PAYMENTS
    var paymentsSum = 0;
    payments?.forEach(p => {
        paymentsSum += p.Value;
    });
    // payment table
    const tableSimple: TableSource = {
        // A list of heading labels.
        head: ['Date', 'Value'],
        // The data visibly shown in your table body UI.
        body: tableMapperValues(payments!, ['Date', 'Value']),
        // Optional: The data returned when interactive is enabled and a row is clicked.
        meta: tableMapperValues(payments!, ['ID']),
        // Optional: A list of footer labels.
        foot: ['Total', '<code class="code">'+paymentsSum+'</code>']
    };

    let selected;

    function mySelectionHandler(meta: unknown): void{
        console.log('on:selected', meta);
        window.location.href = "/expenses/" + meta.detail[0];
    }
</script>

<style>
    .icon {
      width: 32px;
      height: 32px;
    }
  </style>

{#if expense}
<h1 class="h2">{expense.Title}</h1>
<span>Created: {expense.Created_date} </span>

<div class="w-full text-token grid grid-cols-1 md:grid-cols-2 gap-4">
    <div class="card p-4 w-full space-y-4 variant-soft">
        <TabGroup>
            <Tab bind:group={tabSet} name="tab1" value={0}>
                <svelte:fragment slot="lead"><div class="flex justify-center items-center"><div class="icon"><MdEdit/></div></div></svelte:fragment>
                <span>Expense</span>
            </Tab>
            <Tab bind:group={tabSet} name="tab2" value={1}>
                <svelte:fragment slot="lead"><div class="flex justify-center items-center"><div class="icon"><MdPayment/></div></div></svelte:fragment>
                <span>Payments</span>
            </Tab>
            <!-- Tab Panels --->
            <svelte:fragment slot="panel">
                {#if tabSet === 0}
                    <label class="label m-4">
                        <span>Date</span>
                        <input class="input" type="date" value={expense.Date} bind:this={dateInput}/>
                    </label>
                    <label class="label m-4">
                        <span>Price</span>
                        <input class="input" type="number" min="0" step="0.01" value={expense.Value} bind:this={priceInput}/>
                    </label>
                    <label class="label m-4">
                        <span>Correspondent</span>
                        <select class="select" value={expense.Correspondent}>
                            {#each correspondents as correspondent}
                                <option value={correspondent.id}>{correspondent.name}</option>
                            {/each}
                        </select>
                    </label>
                    <div class="flex justify-center items-center m-4">
                        <button type="button" class="btn-lg variant-filled" on:click={SendForm}>Update</button>
                    </div>
                {:else if tabSet === 1}
                    <Table source={tableSimple} interactive={true} on:selected={mySelectionHandler} />
                {/if}
            </svelte:fragment>
        </TabGroup>        
    </div>
    <!-- Thumbnail -->
    <div class="card p-4 w-full text-token space-y-4">
        <a href="https://docs.local/api/documents/{expense.PaperlessID}/preview/" target="_blank">
            <img class="h-auto max-w-full rounded-lg" src="http://localhost:8080/expenses/{expense.PaperlessID}/thumb/" alt={expense.Content}>
        </a>
    </div>
</div>
{:else}
<div class="card p-4 w-full space-y-4">
    Error when retrieving data: {error}
</div>
{/if}