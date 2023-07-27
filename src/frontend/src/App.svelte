<!-- App.svelte -->
<script>
    import { Router, Link, Route } from "svelte-routing";
    import ExpenseOverview from "./ExpenseOverview.svelte";
    import ExpenseDetails from "./ExpenseDetails.svelte";
    export let url = "";
</script>
<Router url="{url}">
  <nav class="navbar">
    <Link to="/">Home</Link>
  </nav>
  <div class="main-container">
    <Route path="movie/:slug" component="{ExpenseDetails}" />
    <Route path="/"><ExpenseOverview /></Route>
  </div>
</Router>
<style>
    .navbar {
        background-color:rgba(0,0,0,0.6);
        display: flex;
        padding: 16px 64px;
        font-weight: bold;
        color:white;
    }
    .main-container {
        margin-top:32px;
        display:flex;
        justify-content: center;
        align-items: center;
        background-color: rgba(0,0,0, 0.15);
    }
</style>