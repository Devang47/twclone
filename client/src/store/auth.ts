import { writable } from 'svelte/store';

export const isAuthenticated = writable<boolean>(false);
export const user = writable<User | undefined>();
export const popupOpen = writable<boolean>(false);
export const error = writable<any>();
