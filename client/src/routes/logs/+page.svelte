<script lang="ts">
	import { Search, Download, Filter, ChevronDown, LogIn, LogOut, Plus } from 'lucide-svelte';

	const loadLogsData = () => new Promise(resolve => setTimeout(resolve, 500));

	const logs = [
		{ name: 'Sarah Jenkins', image: 'https://placehold.co/80x40/333/fff?text=XYZ+8992', plate: 'CA-9X3T2', rfid: 'RFID-7712', uhf: 'UHF-0988', status: 'IN', description: 'Authorized Entry', timestamp: 'Oct 24, 2023 • 14:32:01' },
		{ name: 'Marcus Chen', image: 'https://placehold.co/80x40/333/fff?text=LMN+4410', plate: 'NY-77B12', rfid: 'RFID-4421', uhf: 'UHF-3310', status: 'OUT', description: 'Regular Exit', timestamp: 'Oct 24, 2023 • 14:15:44' },
		{ name: 'Elena Rodriguez', image: 'https://placehold.co/80x40/333/fff?text=ABC+1234', plate: 'TX-LMNOP', rfid: 'RFID-8890', uhf: 'UHF-1122', status: 'IN', description: 'VIP Entry', timestamp: 'Oct 24, 2023 • 13:58:12' },
		{ name: 'David Kim', image: 'https://placehold.co/80x40/333/fff?text=XYZ+8992', plate: 'WA-552XX', rfid: 'RFID-1029', uhf: 'UHF-7754', status: 'OUT', description: 'Late Departure', timestamp: 'Oct 24, 2023 • 13:45:00' }
	];
	const statusOptions = ['All', 'In', 'Out'];

	let statusFilter = $state('All');
	let showDropdown = $state(false);
</script>

<div class="mx-auto max-w-8xl flex-1 flex flex-col pt-2 pb-10">
	<div class="flex items-center justify-between mb-8">
		<h1 class="text-2xl font-semibold text-foreground tracking-tight">Activity Logs</h1>
		<button class="cursor-pointer bg-[#2563EB] text-white px-4 py-2.5 rounded-md font-medium text-sm flex items-center gap-2 hover:bg-blue-700 transition-colors shadow-sm focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2">
			<Plus class="w-4 h-4" /> Add Logs
		</button>
	</div>

	{#await loadLogsData()}
		<div class="animate-pulse flex flex-col w-full">
			<div class="h-20 bg-muted/40 rounded-xl border border-border mb-6"></div>
			<div class="h-16 bg-muted/40 rounded-t-xl border border-border"></div>
			<div class="h-[400px] bg-muted/30 border-x border-border"></div>
			<div class="h-16 bg-muted/40 rounded-b-xl border border-border"></div>
		</div>
	{:then}
		<div class="flex flex-col w-full">
			<!-- Toolbar -->
			<div class="flex flex-col xl:flex-row justify-between items-start xl:items-center p-4 border border-border rounded-xl bg-card gap-4 mb-6 shadow-sm">
				<div class="flex flex-col md:flex-row flex-wrap items-center gap-4 w-full xl:w-auto">
					<!-- search Input -->
					<div class="relative w-full md:w-full lg:w-[450px]">
						<Search class="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" />
						<input type="text" placeholder="Search by name, plate, RFID or UHF..." class="pl-9 pr-4 py-2 bg-secondary/10 border border-transparent rounded-md w-full text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors placeholder:text-muted-foreground" />
					</div>
					
					<!-- Status Dropdown -->
					<div class="relative w-full sm:w-auto">
						<button 
							class="flex items-center justify-between w-full sm:w-[140px] px-4 py-2 border border-border rounded-md text-sm font-medium bg-background hover:bg-muted transition-colors cursor-pointer"
							onclick={() => showDropdown = !showDropdown}
							onblur={() => setTimeout(() => showDropdown = false, 150)}
						>
							Status: {statusFilter} <ChevronDown class="w-4 h-4 ml-2 text-muted-foreground" />
						</button>
						{#if showDropdown}
							<div class="absolute z-10 w-full mt-1 bg-background border border-border rounded-md shadow-lg py-1">
								{#each statusOptions as option}
									<button 
										class="block w-full text-left px-4 py-2 text-sm hover:bg-muted transition-colors cursor-pointer {statusFilter === option ? 'text-primary font-medium bg-muted/50' : 'text-foreground'}"
										onclick={() => { statusFilter = option; showDropdown = false; }}
									>
										{option}
									</button>
								{/each}
							</div>
						{/if}
					</div>
					
					<!-- Time Range Tabs -->
					<div class="flex items-center bg-secondary/10 p-1 rounded-lg w-full sm:w-auto">
						<button class="flex-1 sm:flex-none px-4 py-1.5 text-sm font-medium rounded-md bg-white shadow-sm text-foreground">Today</button>
						<button class="flex-1 sm:flex-none px-4 py-1.5 text-sm font-medium rounded-md text-muted-foreground hover:text-foreground transition-colors cursor-pointer">Week</button>
						<button class="flex-1 sm:flex-none px-4 py-1.5 text-sm font-medium rounded-md text-muted-foreground hover:text-foreground transition-colors cursor-pointer">Month</button>
					</div>
				</div>
				
				<!-- Export -->
				<button class="flex items-center justify-center gap-2 w-full xl:w-auto px-4 py-2 border border-border rounded-md text-sm font-medium bg-background hover:bg-muted transition-colors cursor-pointer">
					<Download class="w-4 h-4 text-muted-foreground" /> Export CSV
				</button>
			</div>

			<!-- Table -->
			<div class="overflow-x-auto border border-border rounded-t-xl bg-card shadow-sm">
				<table class="w-full text-sm text-left whitespace-nowrap">
					<thead class="text-xs text-muted-foreground uppercase border-b border-border bg-card/50">
						<tr>
							<th class="px-6 py-5 font-semibold tracking-wider">NO.</th>
							<th class="px-6 py-5 font-semibold tracking-wider">PLATE IMAGE</th>
							<th class="px-6 py-5 font-semibold tracking-wider">NAME</th>
							<th class="px-6 py-5 font-semibold tracking-wider">LICENSE</th>
							<th class="px-6 py-5 font-semibold tracking-wider">RFID</th>
							<th class="px-6 py-5 font-semibold tracking-wider">UHF</th>
							<th class="px-6 py-5 font-semibold tracking-wider">STATUS</th>
							<th class="px-6 py-5 font-semibold tracking-wider">DESCRIPTION</th>
							<th class="px-6 py-5 font-semibold tracking-wider text-right">TIMESTAMP</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						{#each logs as log, i}
							<tr class="hover:bg-muted/30 transition-colors group">
								<td class="px-6 py-4 text-muted-foreground font-medium">{String(i + 1).padStart(2, '0')}</td>
								<td class="px-6 py-4">
									<div class="h-12 w-20 overflow-hidden rounded bg-black">
										<!-- plate image -->
										<img src={log.image} alt={log.plate} class="h-full w-full object-cover grayscale" />
									</div>
								</td>
								<td class="px-6 py-4 font-semibold text-foreground">{log.name}</td>
								<td class="px-6 py-4">
									<span class="px-2.5 py-1.5 rounded bg-gray-300 border border-gray-400 font-medium text-foreground tracking-widest text-[11px] font-mono">{log.plate}</span>
								</td>
								<td class="px-6 py-4 text-muted-foreground font-medium tracking-wide">{log.rfid}</td>
								<td class="px-6 py-4 text-muted-foreground font-medium tracking-wide">{log.uhf}</td>
								<td class="px-6 py-4">
									{#if log.status === 'IN'}
										<span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded border border-emerald-200 bg-emerald-50 text-emerald-600 font-bold text-[10px] uppercase tracking-wider">
											<LogIn class="w-3.5 h-3.5" /> IN
										</span>
									{:else}
										<span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded border border-[#BC4800]/20 bg-[#BC4800]/10 text-[#BC4800] font-bold text-[10px] uppercase tracking-wider">
											<LogOut class="w-3.5 h-3.5" /> OUT
										</span>
									{/if}
								</td>
								<td class="px-6 py-4 text-muted-foreground">{log.description}</td>
								<td class="px-6 py-4 text-muted-foreground font-medium tracking-wide text-right">{log.timestamp}</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>

			<!-- Pagination -->
			<div class="flex flex-col sm:flex-row items-center justify-between p-4 border border-t-0 border-border rounded-b-xl bg-card gap-4 shadow-sm">
				<span class="text-sm text-muted-foreground font-medium">Showing 1 to 4 of 128 entries</span>
				<div class="flex items-center space-x-1.5">
					<button class="px-3 py-1.5 text-sm border border-transparent text-muted-foreground hover:bg-muted rounded-md disabled:opacity-50 transition-colors font-medium">Prev</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-primary bg-[#2563EB] text-white rounded-md font-semibold">1</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-transparent text-foreground hover:bg-muted hover:border-border rounded-md font-medium transition-colors">2</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-transparent text-foreground hover:bg-muted hover:border-border rounded-md font-medium transition-colors">3</button>
					<button class="px-3 py-1.5 text-sm border border-transparent text-foreground hover:bg-muted rounded-md transition-colors font-medium text-muted-foreground hover:text-foreground">Next</button>
				</div>
			</div>
		</div>
	{/await}
</div>
