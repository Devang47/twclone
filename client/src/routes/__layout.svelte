<script lang="ts">
	import ProgressBar from '$lib/components/ProgressBar.svelte';
	import { onMount } from 'svelte';

	import '$styles/global.scss';

	let isLoading = true;
	onMount(() => {
		isLoading = false;
		addObserver();
	});

	const addObserver = () => {
		let observer = new IntersectionObserver(isElScrolledIntoView, {
			root: null,
			rootMargin: '0px',
			threshold: 0.2
		});
		const elements = document.querySelectorAll('.aos');
		elements.forEach((element: any) => {
			observer.observe(element);
			element.style.opacity = '0%';
			element.style.transform = 'translateY(10%)';
		});
		function isElScrolledIntoView(entries: any[]) {
			entries.forEach((entry) => {
				if (entry.isIntersecting) {
					entry.target.classList.add('scrolled-into-view');
				}
			});
		}
	};
</script>

<svelte:head>
	<link rel="icon" type="image/svg+xml" href="/favicon/favicon.svg" />
	<link rel="icon" type="image/png" href="/favicon/favicon.png" />
</svelte:head>

<ProgressBar />
<slot />
