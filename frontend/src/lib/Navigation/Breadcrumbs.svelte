<script lang="ts">
	import { page } from "$app/stores";
	import type { Page } from "@sveltejs/kit";
	import { writable } from "svelte/store";

    let pathVal = "";
    export let path = writable("");
    /*path.subscribe((value: string) => {
        pathVal = value;
    });*/

    var crumbs: {
        label: string,
        href: string,
    }[];

    page.subscribe((value: any) => {
        pathVal = value.url.pathname;
        const tokens = pathVal.split('/').filter((t: string) => t !== '');

        console.log(`${path}`);

        let tokenPath = '';
        crumbs = tokens.map((t: string) => {
            tokenPath += '/' + t;
            return {
                label: t,
                href: tokenPath,
            };
        });

        // Add a way to get home too.
        crumbs.unshift({ label: 'home', href: '/' });
    });

</script>

<ol class="breadcrumb">
    {#if crumbs.length > 1}
        {#each crumbs as c, i}
            {#if i == crumbs.length-1}
                <li class="crumb">{c.label}</li>
            {:else}
                <li class="crumb"><a class="anchor" href={c.href}>{c.label}</a></li>
                <li class="crumb-separator" aria-hidden>&rsaquo;</li>
            {/if}
        {/each}
    {/if}
</ol>