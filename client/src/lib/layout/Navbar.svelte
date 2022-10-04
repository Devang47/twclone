<script lang="ts">
	import { goto } from '$app/navigation';

	import Logo from '$lib/icons/Logo.svelte';
	import { isAuthenticated, user } from '$store/auth';
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
	<div class="logo cursor-pointer" on:click={() => goto('/')}>
		<Logo />
	</div>

	{#if $user?.username}
		<div class="username">
			{$user?.username}
		</div>
	{/if}

	{#if !$isAuthenticated}
		<button class="logout" on:click={() => goto('/login')}> Login </button>
	{/if}
	{#if $isAuthenticated}
		<button class="logout" on:click={logOut}> Logout </button>
	{/if}
</nav>
