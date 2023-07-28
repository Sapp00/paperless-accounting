/** @type {import('./$types').PageLoad} */
import {API_URL} from '$env/static/private';

export async function load({}) {
    const res = await fetch(`${API_URL}/expenses`);
    const item = await res.json();

    return { item };
}