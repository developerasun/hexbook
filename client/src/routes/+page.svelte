<script lang="ts">
	import { PUBLIC_WS_ENDPOINT } from '$env/static/public';
	import { Button } from '$lib/components/ui/button/index.js';
	import type { PageProps } from './$types';
	let count = $state(0);
	let payload = $state('');

	$effect(() => {
		console.log('mounted');
		const socket = new WebSocket(PUBLIC_WS_ENDPOINT);
		socket.onopen = (e) => socket.send('Hello Server!');
		socket.onmessage = (e) => {
			payload = e.data;
		};
		return () => {
			console.info('cleaned');
			socket.close();
		};
	});
	const onCount = () => count++;
	let { data }: PageProps = $props();
</script>

<Button class="my-4 border-2 border-blue-400" onclick={onCount}>Click me</Button>
<p>count: {count}</p>
<p>fetched data: {data.message.message ?? 'none'}</p>
<p>{payload}</p>
