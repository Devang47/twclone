<script lang="ts">
	import { onMount } from 'svelte';
	import { isAuthenticated, user } from '$store/auth';
	import Loader from '$lib/components/Loader.svelte';
	import { goto } from '$app/navigation';
	import auth from '$utils/authService';
	import { page } from '$app/stores';
	import Navbar from '$lib/layout/Navbar.svelte';
	import { getUser } from '$utils/api/user';
	// import CryptoJS from 'crypto-js';
	import { getTweets } from '$utils/api/tweets';
	import { tweetsData } from '$store';

	let loading = true;
	onMount(async () => {
		if (!$user?.email) {
			let auth0Client = await auth.createClient();
			isAuthenticated.set(await auth0Client.isAuthenticated());
			const data = await auth0Client.getUser();

			if (data?.email) {
				const userdata = await getUser(data.email);
				if (!userdata) goto('/login', { replaceState: true });
				user.set(userdata[0]);
			}

			if (!$isAuthenticated) {
				goto(`/login?next=${location.pathname}`);
			}
		}

		let authKey = $user?.uid as string;
		const res = await getTweets(authKey);
		$tweetsData = res.data;

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
