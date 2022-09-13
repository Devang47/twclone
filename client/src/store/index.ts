import { writable } from 'svelte/store';

export let tweetsData = writable<Tweet[] | undefined>();
