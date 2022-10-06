<script lang="ts">
	import { goto } from '$app/navigation';
	import { fade, scale, slide } from 'svelte/transition';
	import { clickOutside } from '$utils/clickOutside';

	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';
	import { handleLikeTweet } from '$utils/tweet';
	import { user } from '$store/auth';
	import { createEventDispatcher, onMount } from 'svelte';
	import { loading, tweetsData } from '$store';
	import { deleteTweet } from '$utils/api/tweets';

	export let data: Tweet;

	let eventDispatcher = createEventDispatcher();
	let optionsMenuOpen = false;
	let likedByUser: boolean;
	onMount(() => {
		likedByUser = Boolean(data.likes.liked_by.find((e) => e === $user?.username));
	});

	const likeTweet = async () => {
		$loading = true;
		let authKey = $user?.uid as string;
		await handleLikeTweet(data.id, authKey);
		likedByUser = !likedByUser;
		$loading = false;
	};

	const handleDeleteTweet = async () => {
		let authKey = $user?.uid as string;
		try {
			const res = await deleteTweet(authKey, data);
			eventDispatcher('tweet-deleted');
		} catch (error) {
			console.log(error);
		}
	};
</script>

<div class="tweet w-full" in:scale={{ start: 0.9, opacity: 0.8 }} out:fade={{ duration: 300 }}>
	<div class="tweet-header">
		<h2
			class="author w-fit"
			on:click={() => {
				if (window.location.pathname === '/user/' + data.author_uid) return;
				tweetsData.set([]);
				loading.set(true);
				goto('/user/' + data.author_uid, { replaceState: true });
			}}
		>
			{data.author}
		</h2>
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
	</div>

	{#if data.author_uid === $user?.uid}
		<button class="options" on:click|stopPropagation={() => (optionsMenuOpen = !optionsMenuOpen)}>
			{#each new Array(3) as _}
				<span />
			{/each}

			{#if optionsMenuOpen}
				<div
					use:clickOutside
					on:click_outside={() => {
						optionsMenuOpen = !optionsMenuOpen;
					}}
					class="options-menu"
					transition:slide={{ duration: 200 }}
				>
					<button on:click|self={handleDeleteTweet}>Delete</button>
				</div>
			{/if}
		</button>
	{/if}
</div>
