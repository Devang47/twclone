<script lang="ts">
	import Logo from '$lib/icons/Logo.svelte';
	import { isAuthenticated } from '$store/auth';
	import auth from '$utils/authService';
	import type { Auth0Client } from '@auth0/auth0-spa-js';
	import { onMount } from 'svelte';

	let auth0client: Auth0Client;

	onMount(async () => {
		auth0client = await auth.createClient();
	});

	const logOut = async () => {
		try {
			await auth0client.logout();
		} catch (e) {
			console.error(e);
		}
	};
</script>

<nav>
	<div class="logo">
		<Logo />
	</div>

	{#if $isAuthenticated}
		<button class="logout" on:click={logOut}> Logout </button>
	{/if}
</nav>
