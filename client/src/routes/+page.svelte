<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import type { PageProps } from './$types';
	let count = $state(0);
	let payload = $state('');

	$effect(() => {
		console.log('mounted');
		const socket = new WebSocket('ws://localhost:3010/ws');
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

	// socket.addEventListener('message', function (event) {
	// 	console.log('Message from server ', event.data);
	// 	payload = event.data;
	// });
</script>

<Button class="my-4 border-2 border-blue-400" onclick={onCount}>Click me</Button>
<p>count: {count}</p>
<p>fetched data: {data.message.message ?? 'none'}</p>
<p>{payload}</p>
