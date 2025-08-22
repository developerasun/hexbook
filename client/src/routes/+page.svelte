<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import type { PageProps } from './$types';
	let count = $state(0);
	let payload = $state('');

	$effect(() => {
		console.log('mounted');
		return () => console.info('cleaned');
	});
	const onCount = () => count++;
	let { data }: PageProps = $props();

	const socket = new WebSocket('ws://localhost:3010/ws');
	socket.addEventListener('open', function (event) {
		socket.send('Hello Server!');
	});

	socket.addEventListener('message', function (event) {
		console.log('Message from server ', event.data);
		payload = event.data;
	});
</script>

<Button class="my-4 border-2 border-blue-400" onclick={onCount}>Click me</Button>
<p>count: {count}</p>
<p>fetched data: {data.message.message ?? 'none'}</p>
<p>{payload}</p>
