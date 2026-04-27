<script lang="ts">
	import StatsGrid from '$lib/components/dashboard/StatsGrid.svelte';
	import VehicleFlow from '$lib/components/dashboard/VehicleFlow.svelte';
	import RecentScans from '$lib/components/dashboard/RecentScans.svelte';
	import { Download, Plus } from 'lucide-svelte';

	// delay for skeleton loading state
	const loadDashboardData = () => new Promise(resolve => setTimeout(resolve, 500));
</script>

<div class="mx-auto max-w-7xl flex-1 flex flex-col pt-2 pb-10">
	<div class="flex flex-col sm:flex-row items-start sm:items-end justify-between mb-8 gap-4">
		<div>
			<h1 class="text-2xl font-semibold text-foreground tracking-tight">Dashboard LPR Parking</h1>
			<p class="text-sm text-muted-foreground mt-1 text-[13px]">Real-time monitoring for Main Campus Garage</p>
		</div>
		<div class="flex items-center gap-2 sm:gap-3 w-full sm:w-auto">
			<button
				class="flex flex-1 sm:flex-none justify-center items-center gap-2 rounded-md border border-border bg-background px-4 py-2 text-sm font-medium text-foreground hover:bg-muted shadow-sm transition-colors"
			>
				<Download class="h-4 w-4" />
				Export
			</button>
			<button
				class="flex flex-1 sm:flex-none justify-center items-center gap-2 rounded-md bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90 shadow-sm transition-colors"
			>
				<Plus class="h-4 w-4" />
					Entry
			</button>
		</div>
	</div>

	{#await loadDashboardData()}
		<!-- Skeleton Loading State -->
		<div class="flex flex-col gap-6 w-full animate-pulse">
			<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6">
				{#each Array(3) as _}
					<div class="h-[128px] rounded-xl border border-border bg-muted shadow-sm"></div>
				{/each}
			</div>

			<div class="grid grid-cols-1 lg:grid-cols-3 gap-6 lg:h-[500px]">
				<div class="lg:col-span-2 rounded-xl border border-border bg-muted shadow-sm min-h-[300px]"></div>
				<div class="lg:col-span-1 rounded-xl border border-border bg-muted shadow-sm min-h-[300px]"></div>
			</div>
		</div>
	{:then}
		<div class="flex flex-col gap-6">
			<StatsGrid />

			<div class="grid grid-cols-1 lg:grid-cols-3 gap-6 lg:h-[500px]">
				<div class="lg:col-span-2">
					<VehicleFlow />
				</div>
				<div class="lg:col-span-1">
					<RecentScans />
				</div>
			</div>
		</div>
	{/await}
</div>
