<script lang="ts">
	import LoggedInOnly from '$lib/components/LoggedInOnly.svelte';
	import Tweet from '$lib/layout/Tweet.svelte';
	import ArrowBottom from '$lib/icons/ArrowBottom.svelte';
	import TweetsBg from '$lib/layout/TweetsBG.svelte';
	import TweetInput from '$lib/components/TweetInput.svelte';
	import { onMount } from 'svelte';

	import axios from 'axios';

	onMount(async () => {
		const res = await axios.get('http://127.0.0.1:5000/get-tweets');
		console.log(res);
	});
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
			{#each new Array(10) as _}
				<Tweet data={{}} />
			{/each}
		</div>

		<button> Load more <ArrowBottom /> </button>
	</section>
</LoggedInOnly>
