<script lang="ts">
	import { RotateCcw } from 'lucide-svelte';

	interface Scan {
		id: string;
		plate: string;
		type: 'IN' | 'OUT';
		time: string;
		gate: string;
		image: string;
	}

	let scans = $state<Scan[]>([
		{
			id: '1',
			plate: 'XYZ - 8992',
			type: 'IN',
			time: 'Just now',
			gate: 'Gate A',
			image: 'https://placehold.co/80x40/333/fff?text=XYZ+8992'
		},
		{
			id: '2',
			plate: 'LMN - 4410',
			type: 'OUT',
			time: '2 min ago',
			gate: 'Gate C',
			image: 'https://placehold.co/80x40/333/fff?text=LMN+4410'
		},
		{
			id: '3',
			plate: 'ABC - 1234',
			type: 'IN',
			time: '5 min ago',
			gate: 'Gate B',
			image: 'https://placehold.co/80x40/333/fff?text=ABC+1234'
		}
	]);
</script>

<div class="flex h-full flex-col rounded-xl border border-border bg-card shadow-sm">
	<div class="flex items-center justify-between border-b border-border p-6">
		<div class="flex flex-col">
			<h3 class="text-sm font-semibold text-foreground">Recent Scans</h3>
			<span class="text-xs text-muted-foreground">Live LPR feed</span>
		</div>
		<button class="rounded-md p-2 hover:bg-muted text-primary transition-colors">
			<RotateCcw class="h-4 w-4" />
		</button>
	</div>

	<div class="flex-1 overflow-y-auto p-6 space-y-6">
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
