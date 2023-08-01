var margin = {top: 60, right: 230, bottom: 50, left: 50},
    width = 660 - margin.left - margin.right,
    height = 400 - margin.top - margin.bottom;

var svg = d3.select("#expenses_chart")
    .attr("width", width + margin.left + margin.right)
    .attr("height", height + margin.top + margin.bottom)
    .append("g")
    .attr("transform",
        "translate(" + margin.left + "," + margin.top + ")");

var data = JSON.parse(document.currentScript.nextElementSibling.textContent);

console.log(data);

var keys = "year";

var color = d3.scaleOrdinal()
    .domain(keys)
    .range(d3.schemeSet2);


//////////
// AXIS //
//////////

// Add X axis
const timeParse = d3.timeParse("%Y-%b-%d");
var x = d3.scaleTime()
    .domain([new Date(2022, 0, 1), Date.now()])
    .range([ 0, width ]);
var xAxis = svg.append("g")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(x).ticks(5));

// Add X axis label:
svg.append("text")
    .attr("text-anchor", "end")
    .attr("x", width)
    .attr("y", height+40 )
    .text("Time (date)");

// Add Y axis label:
svg.append("text")
    .attr("text-anchor", "end")
    .attr("x", 0)
    .attr("y", -20 )
    .text("Euro")
    .attr("text-anchor", "start")

// Add Y axis
var y = d3.scaleLinear()
    .domain(d3.extent(data, function(d) {return d.value}))
    .range([ height, 0 ]);
svg.append("g")
    .call(d3.axisLeft(y).ticks(5))

// draw the lines
const sumstat = d3.group(data, d => d.category);
const line = d3.line()
  .x(d => x(d.date))
  .y(d => y(d.value));

console.log(sumstat);

//svg.selectAll("path.line")
//    .data(sumstat)
//    .join("path")
svg.selectAll(".line")
.data(sumstat)
.join("path")
 // .attr("fill", "none")
  .attr("fill", function(d){ return color(d[0]) })
  .attr("stroke-width", 1.5)
  .attr("d", function(d){
    return d3.area()
      .x(function(d) { return x(new Date(d.date)); })
      .y0(function(d) { return y(d.value); })
      .y1(function(d) { return height; })
      (d[1])
  });

console.log("yo")