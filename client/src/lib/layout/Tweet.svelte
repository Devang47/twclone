<script lang="ts">
	import { goto } from '$app/navigation';
	import { scale } from 'svelte/transition';

	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';
	import { handleLikeTweet } from '$utils/tweet';
	import { user } from '$store/auth';
	import { onMount } from 'svelte';
	import { loading, tweetsData } from '$store';

	export let data: Tweet;

	let likedByUser: boolean;
	onMount(() => {
		likedByUser = Boolean(data.likes.liked_by.find((e) => e === $user?.username));
		console.log({ likedByUser });
	});

	const likeTweet = async () => {
		$loading = true;
		let authKey = $user?.uid as string;
		await handleLikeTweet(data.id, authKey);
		likedByUser = !likedByUser;
		$loading = false;
	};
</script>

<!-- {
    "_id": "630f6d8d7274a9291225d764",
    "author": "devang",
    "likes": {
        "likedby": null,
        "totallikes": 0
    },
    "publishedon": "2022-08-31 19:47:49.408663 +0530 IST m=+1.676021168",
    "tweet": "deva2314sdfsjhgsdfd2ng@ag.com"
} -->

<div class="tweet w-full" transition:scale={{ start: 0.9, opacity: 0.8 }}>
	<div class="tweet-header">
		<h2 class="author w-fit" on:click={() => goto('/')}>{data.author}</h2>
		<div class="published-date">
			{new Date(data.published_on).toLocaleDateString('en-US', {
				year: 'numeric',
				month: 'short',
				day: 'numeric'
			})}
		</div>
	</div>

	<p>
		{data.tweet}
	</p>

	<div class="action-btns">
		<button class={likedByUser ? 'likedByUser' : ''} on:click={likeTweet}>
			<Heart />
		</button>
		<button>
			<Retweet />
		</button>
	</div>

	<div class="options">
		{#each new Array(3) as _}
			<span />
		{/each}
	</div>
</div>
