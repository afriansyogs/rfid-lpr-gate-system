<script lang="ts">
	import { RotateCcw } from 'lucide-svelte';

	import dummyData from '$lib/data/dummy.json';
	import type { Scan } from '$lib/types';

	let scans = $state<Scan[]>(dummyData.recentScans as Scan[]);
	let isLoading = $state(false);

	function handleRefresh() {
		isLoading = true;
		setTimeout(() => {
			isLoading = false;
		}, 500);
	}
</script>

<div class="flex h-full flex-col rounded-xl border border-border bg-card shadow-sm">
	<div class="flex items-center justify-between border-b border-border p-6">
		<div class="flex flex-col">
			<h3 class="text-sm font-semibold text-foreground">Recent Scans</h3>
			<span class="text-xs text-muted-foreground">Live LPR feed</span>
		</div>
		<button 
			class="rounded-md p-2 hover:bg-muted text-primary transition-colors cursor-pointer disabled:opacity-50"
			onclick={handleRefresh}
			disabled={isLoading}
		>
			<RotateCcw class="h-4 w-4 {isLoading ? 'animate-spin' : ''}" />
		</button>
	</div>

	<div class="flex-1 overflow-y-auto p-6 space-y-6">
		{#if isLoading}
			{#each [1, 2, 3] as _}
				<div class="flex items-center justify-between gap-4 animate-pulse">
					<div class="flex items-center gap-4">
						<div class="h-12 w-20 rounded bg-muted"></div>
						<div class="flex flex-col gap-2">
							<div class="h-4 w-24 rounded bg-muted"></div>
							<div class="h-3 w-32 rounded bg-muted"></div>
						</div>
					</div>
					<div class="h-6 w-12 rounded bg-muted"></div>
				</div>
			{/each}
		{:else}
			{#each scans as scan (scan.id)}
				<div class="flex items-center justify-between gap-4">
				<div class="flex items-center gap-4">
					<div class="h-12 w-20 overflow-hidden rounded bg-black">
						<!-- plate image -->
						<img src={scan.image} alt={scan.plate} class="h-full w-full object-cover grayscale" />
					</div>
					<div class="flex flex-col">
						<span class="font-bold text-foreground">{scan.plate}</span>
						<span class="flex items-center gap-2 text-xs text-muted-foreground mt-1">
							<span><Clock class="h-3 w-3 inline mb-0.5"/> {scan.time}</span>
							<span class="h-1 w-1 rounded-full bg-neutral"></span>
							<span>{scan.gate}</span>
						</span>
					</div>
				</div>

				<div
					class="rounded-md px-2 py-1 text-xs font-bold leading-none
					{scan.type === 'IN'
						? 'bg-blue-100 text-primary'
						: 'bg-red-100 text-tertiary'}"
				>
					{scan.type}
				</div>
			</div>
		{/each}
		{/if}
	</div>

	<div class="border-t border-border p-4">
		<a href="/logs" class="block w-full text-center text-sm font-medium text-black hover:underline">
			View All Activity
		</a>
	</div>
</div>

<script module>
	import { Clock } from 'lucide-svelte';
</script>
