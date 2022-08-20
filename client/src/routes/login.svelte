<script lang="ts">
	import Logo from '$lib/icons/Logo.svelte';
	import ArrowRight from '$lib/icons/ArrowRight.svelte';
	import BlobBlue from '$lib/icons/BlobBlue.svelte';
	import BlobGreen from '$lib/icons/BlobGreen.svelte';
	import auth from '$utils/authService';
	import { isAuthenticated, user } from '$store/auth';
	import { popupOpen } from '$store/auth';
	import Loader from '$lib/components/Loader.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import type { Auth0Client } from '@auth0/auth0-spa-js';

	let auth0Client: Auth0Client;

	onMount(async () => {
		auth0Client = await auth.createClient();
		isAuthenticated.set(await auth0Client.isAuthenticated());
		user.set(await auth0Client.getUser());

		if ($isAuthenticated) {
			goto('/');
		} else {
			goto('/login');
		}
	});

	function login() {
		auth.loginWithPopup(auth0Client, {});
	}
</script>

<svelte:head>
	<title>Signin | TWClone</title>
</svelte:head>

{#if $popupOpen}
	<Loader />
{:else}
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
{/if}
