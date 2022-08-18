<script>
	import '$styles/styles.scss';
	import '$styles/login.scss';
	import Logo from '$lib/icons/Logo.svelte';
	import ArrowRight from '$lib/icons/ArrowRight.svelte';
	import BlobBlue from '$lib/icons/BlobBlue.svelte';
	import BlobGreen from '$lib/icons/BlobGreen.svelte';
</script>

<main>
	<section>
		<div class="container">
			<div class="blob-1">
				<BlobBlue />
			</div>
			<div class="blob-2">
				<BlobGreen />
			</div>
			<div class="inner-wrapper">
				<div class="">
					<div class="logo">
						<Logo />
					</div>
					<h1>Login into <span> TWCLONE </span></h1>
					<button
						>Login

						<ArrowRight />
					</button>
				</div>
			</div>
		</div>
	</section>
</main>
