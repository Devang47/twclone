<script lang="ts">
	import LoggedInOnly from '$lib/components/LoggedInOnly.svelte';
	import Tweet from '$lib/layout/Tweet.svelte';
	import ArrowBottom from '$lib/icons/ArrowBottom.svelte';
	import TweetsBg from '$lib/layout/TweetsBG.svelte';
	import TweetInput from '$lib/components/TweetInput.svelte';
	import { loading, tweetsData } from '$store';
	import { onMount } from 'svelte';
	import { getTweets } from '$utils/api/tweets';
	import { user } from '$store/auth';

	onMount(() => {
		$loading = false;
	});

	let limit = 15;
	const loadMoreTweets = async () => {
		$loading = true;
		limit += 15;
		let authKey = $user?.uid as string;
		const res = await getTweets(authKey, limit);
		$tweetsData = res.data;
		$loading = false;
	};
</script>

<svelte:head>
	<title>App | TWClone</title>
</svelte:head>

<LoggedInOnly>
	<TweetsBg />
	<section id="app">
		<div class="input">
			<TweetInput />
		</div>

		<div class="tweets">
			{#if $tweetsData?.length}
				{#each $tweetsData as item (item.id)}
					<Tweet data={item} />
				{/each}
			{/if}
		</div>

		<button on:click={loadMoreTweets}> Load more <ArrowBottom /> </button>
	</section>
</LoggedInOnly>
