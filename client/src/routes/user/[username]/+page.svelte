<script lang="ts">
	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';
	import { onMount } from 'svelte';
	import { loading, socket } from '$store';
	import { getUser, getUserByUid } from '$utils/api/user';
	import { page } from '$app/stores';

	import TweetInput from '$lib/components/TweetInput.svelte';
	import { isAuthenticated, user } from '$store/auth';
	import { getTweetsByUser } from '$utils/api/tweets';
	import auth from '$utils/authService';
	import { goto } from '$app/navigation';
	import TweetItem from '$lib/layout/Tweet.svelte';
	import Navbar from '$lib/layout/Navbar.svelte';
	import { apiAddr } from '$utils/api/base';
	import { json } from '@sveltejs/kit';

	let userData: User;
	let error: any = null;

	let tweets: Tweet[] = [];

	let uid = $page.params.username;
	$: uid = $page.params.username;

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

		const uid = $page.params.username;
		try {
			userData = (await getUserByUid(uid)) as User;
			loadTweets();
		} catch (err: any) {
			error = err.msg;
		}

		$socket = new WebSocket(`ws://twclone.awsapp.xyz/api/ws`);

		$socket.onerror = (error) => {
			console.log('Socket Error: ', error);
		};

		$socket.onmessage = (msg) => {
			let decodedData = JSON.parse(msg.data);

			if (decodedData.msg === 'user-tweets-updated' && decodedData.user_id === uid) {
				loadTweets();
			}
		};

		$loading = false;
	});

	const loadTweets = async () => {
		tweets = (await getTweetsByUser(uid)).data;
	};
</script>

<svelte:head>
	<title>App | TWClone</title>
</svelte:head>

{#if error}
	{#if error?.msg === 'not-found'}
		404 Not Found
	{:else if error?.msg === 'unknown-error'}
		Unknown Error
	{/if}
{/if}

<Navbar />

<main>
	<section id="user-data">
		<div class="banner">
			<img src="/images/{Math.floor(Math.random() * 4)}.png" alt="" />
		</div>

		<div class="dp">
			<img
				src={userData?.photo}
				alt={'photo of ' + userData?.username}
				referrerpolicy="no-referrer"
			/>
		</div>

		<div class="user-name">
			{userData?.username}
		</div>

		<div class="body">
			<div class="tweets-wrapper">
				<div class="input">
					<TweetInput
						on:tweet-created={() => {
							$socket.send(JSON.stringify({ msg: 'user-tweets-updated', user_id: uid }));
							loadTweets();
						}}
					/>
				</div>

				<div class="" />

				<div class="tweets">
					{#if tweets}
						{#each tweets as item (item.id)}
							<TweetItem on:tweet-deleted={loadTweets} data={item} />
						{/each}
					{/if}
				</div>
			</div>
		</div>
	</section>
</main>
