<script lang="ts">
    import type { PageData } from './$types';  
    import { onMount } from "svelte";
    import * as d3 from 'd3';
	import type { DataType, ItemType } from './+page';

	export let data: PageData;
    const { item, error }: DataType = data;
</script>

<p> detail for XYZ </p>
{#if item}
<p>{item.Date}</p>
<p>{item.Title}</p>
{/if}