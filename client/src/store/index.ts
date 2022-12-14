import { writable } from 'svelte/store';

export let tweetsData = writable<Tweet[] | undefined>();
export let loading = writable<Boolean>(true);
export let socket = writable<WebSocket>()