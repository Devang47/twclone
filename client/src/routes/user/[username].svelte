<script lang="ts">
	import LoggedInOnly from '$lib/components/LoggedInOnly.svelte';
	import Navbar from '$lib/layout/Navbar.svelte';
	import Heart from '$lib/icons/Heart.svelte';
	import Retweet from '$lib/icons/Retweet.svelte';

	import '$styles/_nav.scss';
	import '$styles/_user.scss';
</script>

<svelte:head>
	<title>App | TWClone</title>
</svelte:head>

<LoggedInOnly>
	<section>
		<div class="banner">
			<img src="/images/{Math.floor(Math.random() * 4)}.png" alt="" />
		</div>

		<div class="dp">
			<img
				src="https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1180&q=80"
				alt=""
			/>
		</div>

		<div class="tweets-wrapper">
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
							Lorem ipsum dolor sit, amet consectetur adipisicing elit. Sit earum non quas suscipit!
							Architecto, neque, provident ut non alias eos vitae quam impedit, quaerat molestias
							laudantium sint voluptas minima nobis!
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
		</div>
	</section>
</LoggedInOnly>
