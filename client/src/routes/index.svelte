<script lang="ts">
	import LoggedInOnly from '$lib/components/LoggedInOnly.svelte';
	import Tweet from '$lib/layout/Tweet.svelte';
	import ArrowBottom from '$lib/icons/ArrowBottom.svelte';
	import TweetsBg from '$lib/layout/TweetsBG.svelte';
	import TweetInput from '$lib/components/TweetInput.svelte';
	import { tweetsData } from '$store';
	import { getTweets, postTweets } from '$utils/api/tweets';
	import { user } from '$store/auth';

	const postTweet = async (e: string) => {
		try {
			let authKey = $user?.uid as string;

			await postTweets(authKey, {
				id: '',
				author: $user?.username as string,
				author_uid: $user?.uid as string,
				tweet: e,
				published_on: '',
				likes: { total_likes: 0, liked_by: [] }
			});

			const res = await getTweets(authKey);
			$tweetsData = res.data;

			return true;
		} catch (error) {
			console.log(error);
			return false;
		}
	};
</script>

<svelte:head>
	<title>App | TWClone</title>
</svelte:head>

<LoggedInOnly>
	<TweetsBg />
	<section id="app">
		<div class="input">
			<TweetInput onSubmit={postTweet} />
		</div>

		<div class="tweets">
			{#if $tweetsData?.length}
				{#each $tweetsData.reverse() as item}
					<Tweet data={item} />
				{/each}
			{/if}
		</div>

		<button> Load more <ArrowBottom /> </button>
	</section>
</LoggedInOnly>
