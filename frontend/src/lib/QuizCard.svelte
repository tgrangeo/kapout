<script lang="ts">
	import Button from "./Button.svelte";
	let{ host, quiz } = $props();
</script>

<div class="flex justify-between items-center bg-white border-2 border-green-500 p-4 rounded-xl">
	{quiz}
	<Button on:click={host}>Host</Button>
</div>