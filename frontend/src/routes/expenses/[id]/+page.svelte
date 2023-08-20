<script lang="ts">
    import type { PageData } from './$types';  
    import { onMount } from "svelte";
	import type { Expense, Payment } from './+page';
    import {isValidFloat2D} from '$lib/helper/validation';
    import EditPayment from '$lib/Views/EditPayment.svelte';

    import type { TableSource, ToastSettings, ModalComponent, ModalSettings } from '@skeletonlabs/skeleton';
    import { Table, tableMapperValues, toastStore, TabGroup, Tab, modalStore } from '@skeletonlabs/skeleton';

	export let data: PageData;
    const { expense, payments, correspondents, error } = data;

    import MdEdit from 'svelte-icons/md/MdEdit.svelte';
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
                    expense!.Value = parseFloat(priceInput.value);
                    expense!.Date = dateInput.value;
                }
            });
        } else {
            priceInput.classList.add("input-error");
            toastStore.trigger(tErr);
        }
    }

    // PAYMENTS
    var paymentsSum = 0;
    function UpdatePayments(){
        paymentsSum = 0;
        payments?.forEach(p => {
            paymentsSum += p.Value;
        });
    };
    UpdatePayments();

    function UpdatePaymentEntry(p: Payment){
        for(let i=0; i < payments.length; i++){
            let pp = payments[i];
            if ( pp.ID == p.ID ){
                payments[i] = p;
                return;
            }
        }
        payments.push(p);
        UpdatePayments();
    }

    // payment table
    $: tableSimple = {
        // A list of heading labels.
        head: ['Date', 'Value'],
        // The data visibly shown in your table body UI.
        body: tableMapperValues(payments!, ['Date', 'Value']),
        // Optional: The data returned when interactive is enabled and a row is clicked.
        meta: tableMapperValues(payments!, ['ID', 'Date', 'Value']),
        // Optional: A list of footer labels.
        foot: ['<b>Total</b>', '<b>'+paymentsSum.toFixed(2)+' / '+expense?.Value+'</b>']
    };

    let selected;

    function mySelectionHandler(meta: any): void{
        console.log('on:selected', meta);
        const updatePaymentModalComponent: ModalComponent = {
            // Pass a reference to your custom component
            ref: EditPayment,
            // Add the component properties as key/value pairs
            props: {  
                head: "Update Payment",
                payment: {ID: meta.detail[0], Date: meta.detail[1], Value: meta.detail[2], ExpenseID: expense!.PaperlessID} as Payment, 
                expense: expense, 
                class: "card p-4"
            },
        };
        const updatePaymentModal: ModalSettings = {
            type: 'component',
            
            // Pass the component directly:
            component: updatePaymentModalComponent,
            response: (r: Payment) => {
                if (r != undefined) {
                    console.log(r);
                    UpdatePaymentEntry(r);
                    modalStore.close(); 
                } else {
                    console.log("just closed");
                }
            }
        };
        modalStore.trigger(updatePaymentModal);
        //window.location.href = "/payments/" + meta.detail[0];
    }
    function AddPayment(){
        const tTodo: ToastSettings = {
            message: 'Not yet implemented'
        };
        toastStore.trigger(tTodo);

        const addPaymentModalComponent: ModalComponent = {
            // Pass a reference to your custom component
            ref: EditPayment,
            // Add the component properties as key/value pairs
            props: {  
                head: "Add Payment",
                payment: null, 
                expense: expense, 
                class: "card p-4"
            },
        };
        const addPaymentModal: ModalSettings = {
            type: 'component',
            
            // Pass the component directly:
            component: addPaymentModalComponent,
            response: (r: Payment) => {
                if(r != undefined){
                    console.log(r);
                    UpdatePaymentEntry(r);
                    modalStore.close();
                } else {
                    console.log("didnt do anything");
                }
            }
        };

        modalStore.trigger(addPaymentModal);
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
                        <button type="button" class="btn variant-filled w-full" on:click={SendForm}>Update</button>
                    </div>
                {:else if tabSet === 1}
                    <div class="flex justify-center items-center m-4">
                        <button type="button" class="btn variant-filled w-full" on:click={AddPayment}>Add Payment</button>
                    </div>
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