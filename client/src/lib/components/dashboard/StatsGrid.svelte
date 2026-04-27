<script lang="ts">
	import { Car, LogIn, LogOut, LayoutGrid, TrendingUp } from 'lucide-svelte';
</script>

<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6">
	<!-- Currently Inside -->
	<div class="rounded-xl border border-border bg-card p-6 shadow-sm">
		<div class="flex items-center justify-between">
			<h3 class="text-sm font-medium text-muted-foreground">Currently Inside</h3>
			<div class="flex h-8 w-8 items-center justify-center rounded-md bg-blue-50">
				<Car class="h-4 w-4 text-primary" />
			</div>
		</div>
		<div class="mt-4">
			<span class="text-3xl font-bold text-foreground">342</span>
			<div class="mt-2 flex items-center text-sm">
				<span class="flex items-center text-green-600 font-medium">
					<TrendingUp class="mr-1 h-3 w-3" />
					+12%
				</span>
				<span class="ml-2 text-muted-foreground">vs last hour</span>
			</div>
		</div>
	</div>

	<!-- Exited Today -->
	<div class="rounded-xl border border-border bg-card p-6 shadow-sm">
		<div class="flex items-center justify-between">
			<h3 class="text-sm font-medium text-muted-foreground">Exited Today</h3>
			<div class="flex h-8 w-8 items-center justify-center rounded-md bg-red-50">
				<LogOut class="h-4 w-4 text-tertiary" />
			</div>
		</div>
		<div class="mt-4">
			<span class="text-3xl font-bold text-foreground">942</span>
			<div class="mt-2 text-sm text-muted-foreground">
				Total departures since 00:00
			</div>
		</div>
	</div>

	<!-- Entered Today -->
	<div class="rounded-xl border border-border bg-card p-6 shadow-sm">
		<div class="flex items-center justify-between">
			<h3 class="text-sm font-medium text-muted-foreground">Entered Today</h3>
			<div class="flex h-8 w-8 items-center justify-center rounded-md bg-green-50">
				<LogIn class="h-4 w-4 text-green-600" />
			</div>
		</div>
		<div class="mt-4">
			<span class="text-3xl font-bold text-foreground">1,284</span>
			<div class="mt-2 text-sm text-muted-foreground">
				Total arrivals since 00:00
			</div>
		</div>
	</div>
</div>
