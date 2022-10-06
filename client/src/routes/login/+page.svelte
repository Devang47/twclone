<script lang="ts">
	import Logo from '$lib/icons/Logo.svelte';
	import ArrowRight from '$lib/icons/ArrowRight.svelte';
	import BlobBlue from '$lib/icons/BlobBlue.svelte';
	import BlobGreen from '$lib/icons/BlobGreen.svelte';
	import auth from '$utils/authService';
	import { isAuthenticated, user } from '$store/auth';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import type { Auth0Client } from '@auth0/auth0-spa-js';
	import { getUser } from '$utils/api/user';
	import { loading } from '$store';

	let auth0Client: Auth0Client;

	onMount(async () => {
		auth0Client = await auth.createClient();
		isAuthenticated.set(await auth0Client.isAuthenticated());

		const data = await auth0Client.getUser();

		if (data?.email) {
			const userdata = await getUser(data.email);
			if (!userdata) return goto('/login', { replaceState: true });

			user.set(userdata[0]);
		}

		$loading = false;

		if ($isAuthenticated) {
			goto($page.url.searchParams.get('next') || '/');
		}
	});

	const login = async () => {
		if (await auth.loginWithPopup(auth0Client, {})) {
			goto($page.url.searchParams.get('next') || '/');
		}
	};
</script>

<svelte:head>
	<title>Signin | TWClone</title>
</svelte:head>

<main>
	<section id="login">
		<div class="container">
			<div class="blob-1">
				<BlobBlue />
			</div>
			<div class="blob-2">
				<BlobGreen />
			</div>
			<div class="inner-wrapper">
				<div class="">
					<div class="logo">
						<Logo />
					</div>
					<h1>Login into <span> TWCLONE </span></h1>
					<button on:click={login}>
						Login
						<ArrowRight />
					</button>
				</div>
			</div>
		</div>
	</section>
</main>
