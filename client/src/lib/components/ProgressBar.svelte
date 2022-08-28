<script lang="ts">
	import { navigating } from '$app/stores';

	let progress = 0;
	let isNavigated = false;
	let progressBar: any;

	$: {
		if ($navigating) {
			progress = 30;
			isNavigated = true;
			progressBar && (progressBar.style.width = '30%');
			progressBar && (progressBar.style.opacity = '100%');
		}
		if (!$navigating && isNavigated) {
			progress = 100;
			progressBar && (progressBar.style.width = '100%');
			isNavigated = false;
		}
	}

	$: {
		if (progress == 100) {
			progress = 0;
			setTimeout(() => {
				progressBar && (progressBar.style.opacity = '0%');
				progressBar && (progressBar.style.width = '0%');
			}, 500);
		}
	}
</script>

<div class="progress-bar" bind:this={progressBar} role="progressbar" />

<style lang="postcss">
	.progress-bar {
		@apply fixed top-0 left-0 h-[2.5px] z-40 bg-complementary;
		transition: width 0.5s ease, opacity 0.3s linear;
	}
</style>
