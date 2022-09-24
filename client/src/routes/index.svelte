<script lang="ts">
	import LoggedInOnly from '$lib/components/LoggedInOnly.svelte';
	import Tweet from '$lib/layout/Tweet.svelte';
	import ArrowBottom from '$lib/icons/ArrowBottom.svelte';
	import TweetsBg from '$lib/layout/TweetsBG.svelte';
	import TweetInput from '$lib/components/TweetInput.svelte';
	import { loading, tweetsData } from '$store';
	import { onMount } from 'svelte';
	import { handlePostTweet } from '$utils/tweet';

	onMount(() => {
		$loading = false;
	});
</script>

<svelte:head>
	<title>App | TWClone</title>
</svelte:head>

<LoggedInOnly>
	<TweetsBg />
	<section id="app">
		<div class="input">
			<TweetInput onSubmit={handlePostTweet} />
		</div>

		<div class="tweets">
			{#if $tweetsData?.length}
				{#each $tweetsData as item (item.id)}
					<Tweet data={item} />
				{/each}
			{/if}
		</div>

		<button> Load more <ArrowBottom /> </button>
	</section>
</LoggedInOnly>
