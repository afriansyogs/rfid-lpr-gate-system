<script lang="ts">
	import { ChevronDown } from 'lucide-svelte';

	let filter = $state('Today');

	const chartData = [
		{ time: '8am', in: 75, out: 25 },
		{ time: '9am', in: 160, out: 55 },
		{ time: '10am', in: 115, out: 85 },
		{ time: '11am', in: 65, out: 95 },
		{ time: '12pm', in: 95, out: 135 },
		{ time: '1pm', in: 75, out: 85 }
	];

	// Map values 0-200 to height percentage 0-100%
	const getPx = (val: number) => `${(val / 200) * 100}%`;
</script>

<div class="flex h-full flex-col rounded-xl border border-border bg-card p-6 shadow-sm">
	<div class="flex items-start justify-between">
		<div>
			<h3 class="text-sm font-semibold text-foreground">Vehicle Flow</h3>
			<p class="text-xs text-muted-foreground mt-1">Hourly IN vs OUT volume</p>
		</div>
		<button
			class="flex items-center gap-2 rounded-md border border-border bg-background px-3 py-1.5 text-sm font-medium hover:bg-muted"
		>
			{filter}
			<ChevronDown class="h-4 w-4 text-muted-foreground" />
		</button>
	</div>

	<!-- Custom Bar Chart Layout -->
	<div class="mt-8 flex flex-1 flex-col">
		<div class="relative flex flex-1">
			<!-- Y Axis -->
			<div class="flex w-8 flex-col justify-between pt-2 pb-6 text-[10px] text-muted-foreground">
				<span>200</span>
				<span>150</span>
				<span>100</span>
				<span>50</span>
				<span>0</span>
			</div>

			<!-- Grid and Chart Area -->
			<div class="relative flex-1 border-b border-border pl-4">
				<!-- Horizontal Grid Lines -->
				<div class="absolute inset-0 flex flex-col justify-between pb-6 pt-2">
					<div class="border-t border-dashed border-border"></div>
					<div class="border-t border-dashed border-border"></div>
					<div class="border-t border-dashed border-border"></div>
					<div class="border-t border-dashed border-border"></div>
					<div class="border-t border-transparent"></div>
				</div>

				<!-- Bars Container -->
				<div class="absolute bottom-6 left-4 right-0 top-2 flex justify-around items-end z-10 px-4">
					{#each chartData as tick}
						<div class="group flex h-full items-end gap-1.5 px-3">
							<!-- Enter Bar -->
							<div
								class="w-4 rounded-t-sm bg-[#2563EB] transition-all hover:bg-[#059669]"
								style="height: {getPx(tick.in)}"
								title="IN: {tick.in}"
							></div>
							<!-- Exit Bar -->
							<div
								class="w-4 rounded-t-sm bg-[#BC4800] transition-all hover:bg-[#d97706]"
								style="height: {getPx(tick.out)}"
								title="OUT: {tick.out}"
							></div>
						</div>
					{/each}
				</div>

				<!-- X Axis -->
				<div class="absolute -bottom-6 left-4 right-0 flex justify-around items-center px-4">
					{#each chartData as tick}
						<span class="text-[10px] text-muted-foreground">{tick.time}</span>
					{/each}
				</div>
			</div>
		</div>

		<!-- Legend -->
		<div class="mt-10 flex items-center justify-center gap-6 pb-2 text-xs text-muted-foreground">
			<div class="flex items-center gap-2">
				<div class="h-3 w-3 rounded-sm bg-primary"></div>
				<span>Entered</span>
			</div>
			<div class="flex items-center gap-2">
				<div class="h-3 w-3 rounded-sm bg-tertiary"></div>
				<span>Exited</span>
			</div>
		</div>
	</div>
</div>
