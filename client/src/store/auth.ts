import type { Auth0Client, User } from '@auth0/auth0-spa-js';
import { writable } from 'svelte/store';

export const isAuthenticated = writable<boolean>(false);
export const user = writable<User | undefined>({});
export const popupOpen = writable<boolean>(false);
export const error = writable<any>();
