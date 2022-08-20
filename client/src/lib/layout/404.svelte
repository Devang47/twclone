<script>
	import { goto } from '$app/navigation';

	import { onMount } from 'svelte';
	let duration = 5;

	onMount(() => {
		setInterval(() => {
			duration > 0 ? (duration -= 1) : goto('/login');
		}, 1000);
	});
</script>

<div class="py-20 text-white">
	<h1>404 - Not Found</h1>

	<h2 class="mt-10">
		redirecting in {duration}
	</h2>
</div>

<style lang="postcss">
	h1 {
		@apply text-center text-[35px] font-bold;
	}

	h2 {
		@apply text-center text-lg;
	}
</style>
