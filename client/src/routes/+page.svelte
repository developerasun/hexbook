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

<Button class="my-4 border-2 border-blue-400" onclick={onCount}>Click me</Button>
<p>count: {count}</p>
<p>fetched data: {data.message.message ?? 'none'}</p>
