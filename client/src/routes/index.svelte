<script>
	import ArrowBottom from '$lib/icons/ArrowBottom.svelte';
	import BlobBlue from '$lib/icons/BlobBlue.svelte';
	import BlobCyan from '$lib/icons/BlobCyan.svelte';
	import BlobGreen from '$lib/icons/BlobGreen.svelte';
	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';
	import Navbar from '$lib/layout/Navbar.svelte';
	import '$styles/styles.scss';
	import '$styles/app.scss';
</script>

<Navbar />
<main>
	<section class="">
		<div class="container">
			<div class="data no-scrollbar">
				<div class="input">
					<div class="textarea">
						<textarea
							name=""
							id=""
							cols="30"
							rows="10"
							maxlength="200"
							class="no-scrollbar"
							placeholder="Write Something..."
						/>
						<div class="cshadow" />
						<button> Tweet </button>
					</div>
				</div>

				<div class="" />

				<div class="tweets">
					{#each new Array(10) as _}
						<div class="tweet">
							<h2 class="author">Devang Saklani</h2>

							<p>
								Lorem ipsum dolor sit, amet consectetur adipisicing elit. Sit earum non quas
								suscipit! Architecto, neque, provident ut non alias eos vitae quam impedit, quaerat
								molestias laudantium sint voluptas minima nobis!
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
					{/each}
				</div>

				<button> Load more <ArrowBottom /> </button>
			</div>
			<div class="blob-1">
				<BlobBlue />
			</div>
			<div class="blob-2">
				<BlobGreen />
			</div>
			<div class="blob-3">
				<BlobCyan />
			</div>
		</div>
	</section>
</main>
