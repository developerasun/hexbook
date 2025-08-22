<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import type { PageProps } from './$types';
	let count = $state(0);
	$effect(() => {
		console.log('mounted');
		return () => console.info('cleaned');
	});
	const onCount = () => count++;
	let { data }: PageProps = $props();
</script>

<h1>Welcome to SvelteKitss</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
<p>{data.message.message}</p>
<!--  -->

<Button class="my-4 border-2 border-blue-400" onclick={onCount}>Click me\</Button>

<p>count: {count}</p>
