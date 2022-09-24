<script lang="ts">
	import { onMount } from 'svelte';

	let tweetInput = '';
	let tweetInputBox: HTMLTextAreaElement;

	export let onSubmit: (e: string) => Promise<boolean>;

	const handleSubmit = () => {
		tweetInput = tweetInput.trim();
		if (!tweetInput) return;

		onSubmit(tweetInput);
		tweetInput = '';
	};

	onMount(() => {
		tweetInputBox.addEventListener('keypress', (e) => {
			if (innerWidth < 640) return;

			if (e.key === 'Enter' && e.shiftKey) {
				e.preventDefault();
				tweetInput += '\r\n';
			} else if (e.key === 'Enter') {
				e.preventDefault();
				handleSubmit();
			}
		});
	});
</script>

<div class="textarea">
	<textarea
		bind:value={tweetInput}
		bind:this={tweetInputBox}
		name="create-new-tweet"
		id="create-tweet"
		cols="30"
		rows="10"
		maxlength="250"
		class="no-scrollbar"
		placeholder="Write Something..."
	/>
	<div class="cshadow" />
	<button on:click={handleSubmit}>
		<svg xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid meet" viewBox="0 0 24 24"
			><path
				fill="currentColor"
				d="m12.815 12.197l-7.532 1.255a.5.5 0 0 0-.386.318L2.3 20.728c-.248.64.421 1.25 1.035.942l18-9a.75.75 0 0 0 0-1.341l-18-9c-.614-.307-1.283.303-1.035.942l2.598 6.958a.5.5 0 0 0 .386.318l7.532 1.255a.2.2 0 0 1 0 .395Z"
			/></svg
		>
	</button>
</div>
