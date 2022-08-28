<script lang="ts">
	import { onMount } from 'svelte';
	import { isAuthenticated, user } from '$store/auth';
	import Loader from '$lib/components/Loader.svelte';
	import { goto } from '$app/navigation';
	import auth from '$utils/authService';
	import { page } from '$app/stores';
	import Navbar from '$lib/layout/Navbar.svelte';

	let loading = true;
	onMount(async () => {
		if (!$user?.email) {
			let auth0Client = await auth.createClient();
			isAuthenticated.set(await auth0Client.isAuthenticated());
			user?.set(await auth0Client.getUser());

			if (!$isAuthenticated) {
				goto(`/login?next=${location.pathname}`);
			}
			console.log($user, auth0Client);
		}

		loading = false;
	});
</script>

{#if loading || !$isAuthenticated}
	<Loader />
{:else}
	<Navbar />
	<main>
		<slot />
	</main>
{/if}
