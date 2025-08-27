<script lang="ts">
	let { text, align }: { text: string; align: 'text-center' | 'text-left' | 'text-right' } =
		$props();
</script>

<p
	class={`leading-7 ${align} w-full max-w-md whitespace-normal break-words [&:not(:first-child)]:mt-6`}
>
	{text}
</p>
