<script lang="ts">
	import './layout.css';
	import Sidebar from '$lib/components/layout/Sidebar.svelte';
	import Header from '$lib/components/layout/Header.svelte';
	
	let { children } = $props();
	let sidebarOpen = $state(false);
	let sidebarCollapsed = $state(false);
</script>

<div class="flex h-screen w-full bg-[#FAFAFA] overflow-hidden relative">
	<Sidebar bind:isOpen={sidebarOpen} bind:isCollapsed={sidebarCollapsed} />
	<div class="flex flex-1 flex-col overflow-hidden min-w-0 h-screen">
		<Header bind:sidebarOpen={sidebarOpen} bind:sidebarCollapsed={sidebarCollapsed} />
		<main class="flex-1 overflow-y-auto p-4 md:p-6 text-foreground h-full block">
			{@render children()}
		</main>
	</div>
</div>
