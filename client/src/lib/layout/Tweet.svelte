<script lang="ts">
	import { goto } from '$app/navigation';
	import { scale } from 'svelte/transition';

	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';

	export let data: Tweet;
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

<div class="tweet w-full" transition:scale={{ duration: 0.3, start: 0.5, opacity: 0.5 }}>
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
		<button>
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
