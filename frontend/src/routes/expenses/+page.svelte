<script lang="ts">
    import type { PageData } from './$types';  
    import { onMount } from "svelte";
    import * as d3 from 'd3';

	export let data: PageData;
    const { items, error } = data;

    type ChartEntry = {
        date: number,
        category: string,
        value: number
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
            .domain([new Date(2022, 0, 1), Date.now()])
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


</script>

<a href="/expenses/detail">for some more details</a>
<svg id="expenses_chart" width="600" height="500"></svg>
<p>script?</p>