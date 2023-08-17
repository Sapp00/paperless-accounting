<script lang="ts">
    import type { PageData } from './$types';  
    import { onMount } from "svelte";
    import * as d3 from 'd3';

    import { Table } from '@skeletonlabs/skeleton';
    import { tableMapperValues, type TableSource } from '@skeletonlabs/skeleton';

	export let data: PageData;
    const { items, expenses, error, fromDate, toDate } = data;

    type ChartEntry = {
        date: String,
        category: String,
        value: Number
    };

    onMount(() => {
        var margin = {top: 20, right: 20, bottom: 30, left: 30},
        width = 928,
        height = 500;

        var svg = d3.select("#expenses_chart")
            .attr("width", width + margin.left + margin.right)
            .attr("height", height + margin.top + margin.bottom)
            .append("g")
            .attr("transform",
                "translate(" + margin.left + "," + margin.top + ")");

        var keys = "date";

        var color = d3.scaleOrdinal()
            .domain(keys)
            .range(d3.schemeSet2);

        // Add X axis
        const timeParse = d3.timeParse("%Y-%b-%d");
        var x = d3.scaleTime()
            .domain([new Date(items[0].date), new Date(items[items!.length-1].date) ])
            .range([ margin.left, width - margin.right ]);
        const xAxis = (g, x) => g 
            .call(d3.axisBottom(x).ticks( width / 80).tickSizeOuter(0));


        const gx = svg.append("g")
            .attr("transform", `translate(0,${height - margin.bottom})`)
            .call(xAxis, x)


        // Add Y axis
        var y = d3.scaleLinear()
            .domain(d3.extent(items, function(d: ChartEntry) {return d.value}))
            .range([ height - margin.bottom, margin.top]);
        var yAxis = svg.append("g")
            .attr("transform", `translate(${margin.left},0)`)
            .call(d3.axisLeft(y).ticks(null, "s"))
            .call( (g: any) => g.select(".domain").remove())
            .call( (g: any) => g.select(".tick:last-of-type text").clone()
                .attr("x", 3)
                .attr("text-anchor", "start")
                .attr("font-weight", "bold")
                .text("Euro"));

        // draw the lines
        const sumstat = d3.group(items, (d: ChartEntry) => d.category);

        // calc area
        const area = (d: any) => d3.area()
                .curve(d3.curveStepAfter)
                .x( (d: ChartEntry) => x(new Date(d.date)) )
                .y0( (d: ChartEntry) => y(d.value) )
                .y1( (d: ChartEntry) => y(0) )
            (d[1]);

        console.log(sumstat);

        //svg.selectAll("path.line")
        //    .data(sumstat)
        //    .join("path")
        const path = svg.selectAll(".line")
            .data(sumstat)
            .join("path")
                .attr("stroke", "none")
                .attr("fill", (d: any) => color(d[0]) )
             .attr("d", (d: any) => area(d));
    });

    
    // table

    let expenseSum = 0;
    let paidSum = 0;
    expenses.forEach(e => {
        expenseSum += e.Value;
        if (e.paidValue != null){
            paidSum += e.paidValue;
        }
    });
    const tableSimple: TableSource = {
        // A list of heading labels.
        head: ['Date', 'Title', 'Price', 'Paid'],
        // The data visibly shown in your table body UI.
        body: tableMapperValues(expenses, ['Date', 'Title', 'Value', 'paidValue']),
        // Optional: The data returned when interactive is enabled and a row is clicked.
        meta: tableMapperValues(expenses, ['PaperlessID', 'Date', 'Title', 'Value']),
        // Optional: A list of footer labels.
        foot: ['<b>Total</b>', '', `<b>${expenseSum}</b>`, `<b>${paidSum}</b>`]
    };


    console.log(expenses);

    let selected;

    function mySelectionHandler(meta: unknown): void{
        console.log('on:selected', meta);
        window.location.href = "/expenses/" + meta.detail[0];
    }

    var dateOptions = { year: 'numeric', month: 'long', day: 'numeric' };

</script>


<h1 class="h3">Expense Overview from {fromDate?.toLocaleDateString("en-US", dateOptions)}-{toDate?.toLocaleDateString("en-US", dateOptions)}</h1>
<svg id="expenses_chart" width="600" height="500"></svg>

<div class="flow-root">
    <div class="btn-group variant-ghost-surface float-left my-2">
        <button>Add Payment</button>
    </div>
    <div class="btn-group variant-ghost-surface float-right my-2">
        <button>Export ZIP</button>
        <button>Export XLSX</button>
        <button>Export Diagram</button>
    </div>
    <!-- better use Paginator -->
    <Table source={tableSimple} interactive={true} on:selected={mySelectionHandler} />
</div>
