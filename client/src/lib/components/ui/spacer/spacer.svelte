<script lang="ts">
  let { my }: { my: 'my-1' | 'my-2' | 'my-3' | 'my-4' } = $props();
</script>

<div class={my}></div>
