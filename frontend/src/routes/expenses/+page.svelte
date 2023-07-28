<script>
    /** @type {import('./$types').PageData} */
   /* import { onMount } from 'svelte';
    import Table from "../../components/Table.svelte";

    type Post = {
        date: Date;
        category: string;
        value: Number;
    };

    let items: Post[] = [];
    let loaded = false;

    onMount(() => loadThings(false))

    function loadThings(wait: boolean) {
            if (typeof fetch !== 'undefined') {
                loaded = false;

                fetch('https://api.fake-rest.refine.dev/posts')
                    .then((response) => response.json())
                    .then((json) =>
                        setTimeout(
                            () => {
                                items = json;
                                loaded = true;
                            },
                            // Simulate a long load time.
                            wait ? 2000 : 0
                        )
                    );
            }
    }*/
    export let data;
</script>

<h1>About this site</h1>
<p>TODO...</p>
<p>{ data.item }</p>
<a href="/">Home</a>