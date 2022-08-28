<script lang="ts">
	import { goto } from '$app/navigation';

	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';

	export let data: any;
</script>

<div class="tweet">
	<h2 class="author" on:click={() => goto('/')}>Devang Saklani</h2>

	<p>
		Lorem ipsum dolor sit, amet consectetur adipisicing elit. Sit earum non quas suscipit!
		Architecto, neque, provident ut non alias eos vitae quam impedit, quaerat molestias laudantium
		sint voluptas minima nobis!
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
