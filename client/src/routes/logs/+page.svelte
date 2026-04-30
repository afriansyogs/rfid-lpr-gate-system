<script lang="ts">
	import { Search, Download, Filter, ChevronDown, LogIn, LogOut, Plus, Loader2, Check, Copy } from 'lucide-svelte';
	import dummyData from '$lib/data/dummy.json';
	import EntryModal from '$lib/components/shared/EntryModal.svelte';
	import { useClipboard } from '$lib/utils/clipboard.svelte';
	import type { ActivityLog, EntryFormData } from '$lib/types';

	const loadLogsData = () => new Promise(resolve => setTimeout(resolve, 500));
	const clipboard = useClipboard();

	let logs = $state<ActivityLog[]>(dummyData.logs as ActivityLog[]);
	const statusOptions = ['All', 'IN', 'OUT'];

	let statusFilter = $state('All');
	let timeFilter = $state('Today');
	let showDropdown = $state(false);
	
	let isEntryModalOpen = $state(false);

	let searchQuery = $state('');
	let isSearching = $state(false);
	let searchTimeout: ReturnType<typeof setTimeout>;

	let filteredLogs = $derived(logs.filter(log => {
		let matchSearch = true;
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			matchSearch = log.name.toLowerCase().includes(q) ||
										log.plate.toLowerCase().includes(q) ||
										log.rfid.toLowerCase().includes(q) ||
										log.uhf.toLowerCase().includes(q);
		}

		let matchStatus = true;
		if (statusFilter !== 'All') {
			matchStatus = log.status.toUpperCase() === statusFilter.toUpperCase();
		}

		return matchSearch && matchStatus;
	}));

	function handleSearchInput(e: Event) {
		const target = e.target as HTMLInputElement;
		const value = target.value;
		isSearching = true;
		
		clearTimeout(searchTimeout);
		searchTimeout = setTimeout(() => {
			searchQuery = value;
			isSearching = false;
		}, 300);
	}

	function handleSaveLog(data: EntryFormData) {
		const now = new Date();
		const timestamp = `${now.toLocaleString('default', { month: 'short' })} ${now.getDate()}, ${now.getFullYear()} • ${now.toLocaleTimeString('en-US', { hour12: false })}`;
		const newLog: ActivityLog = {
			...data,
			status: data.status || 'IN',
			description: data.description || '',
			image: 'https://placehold.co/80x40/333/fff?text=' + encodeURIComponent(data.plate),
			timestamp
		};
		logs = [newLog, ...logs];
		isEntryModalOpen = false;
	}
</script>

<div class="mx-auto max-w-8xl flex-1 flex flex-col pt-2 pb-10">
	<div class="flex items-center justify-between mb-8">
		<h1 class="text-2xl font-semibold text-foreground tracking-tight">Activity Logs</h1>
		<button onclick={() => isEntryModalOpen = true} class="cursor-pointer bg-[#2563EB] text-white px-4 py-2.5 rounded-md font-medium text-sm flex items-center gap-2 hover:bg-blue-700 transition-colors shadow-sm focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2">
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
						<input 
							type="text" 
							oninput={handleSearchInput}
							placeholder="Search by name, plate, RFID or UHF..." 
							class="pl-9 pr-10 py-2 bg-secondary/10 border border-transparent rounded-md w-full text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors placeholder:text-muted-foreground" 
						/>
						{#if isSearching}
							<Loader2 class="absolute right-3 top-2.5 h-4 w-4 text-muted-foreground animate-spin" />
						{/if}
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
						<button onclick={() => timeFilter = 'Today'} class="flex-1 sm:flex-none px-4 py-1.5 text-sm font-medium rounded-md shadow-sm transition-colors cursor-pointer {timeFilter === 'Today' ? 'bg-white text-foreground' : 'text-muted-foreground hover:text-foreground'}">Today</button>
						<button onclick={() => timeFilter = 'Week'} class="flex-1 sm:flex-none px-4 py-1.5 text-sm font-medium rounded-md transition-colors cursor-pointer {timeFilter === 'Week' ? 'bg-white text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}">Week</button>
						<button onclick={() => timeFilter = 'Month'} class="flex-1 sm:flex-none px-4 py-1.5 text-sm font-medium rounded-md transition-colors cursor-pointer {timeFilter === 'Month' ? 'bg-white text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}">Month</button>
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
							<th class="px-6 py-5 font-extrabold tracking-wider">NO.</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">PLATE IMAGE</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">NAME</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">LICENSE</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">RFID ID</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">UHF ID</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">STATUS</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">DESCRIPTION</th>
							<th class="px-6 py-5 font-extrabold tracking-wider text-right">TIMESTAMP</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						{#each filteredLogs as log, i}
							<tr class="hover:bg-muted/30 transition-colors group">
								<td class="px-6 py-4 text-muted-foreground font-medium">{String(i + 1).padStart(2, '0')}</td>
								<td class="px-6 py-4">
									<div class="h-12 w-20 overflow-hidden rounded bg-black">
										<!-- plate image -->
										<img src={log.image} alt={log.plate} class="h-full w-full object-cover grayscale" />
									</div>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(log.name)} title="Copy Name">
										<span class="font-semibold text-foreground">{log.name}</span>
										{#if clipboard.copiedId === log.name}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(log.plate)} title="Copy Plate">
										<span class="px-2.5 py-1.5 rounded bg-gray-300 border border-gray-400 font-medium text-foreground tracking-widest text-[11px] font-mono">{log.plate}</span>
										{#if clipboard.copiedId === log.plate}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(log.rfid)} title="Copy RFID">
										<span>{log.rfid}</span>
										{#if clipboard.copiedId === log.rfid}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(log.uhf)} title="Copy UHF ID">
										<span>{log.uhf}</span>
										{#if clipboard.copiedId === log.uhf}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
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
						{:else}
							<tr>
								<td colspan="9" class="px-6 py-8 text-center text-muted-foreground">
									No activity logs found.
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>

			<!-- Pagination -->
			<div class="flex flex-col sm:flex-row items-center justify-between p-4 border border-t-0 border-border rounded-b-xl bg-card gap-4 shadow-sm">
				<span class="text-sm text-muted-foreground font-medium">Showing {filteredLogs.length} entries</span>
				<div class="flex items-center space-x-1.5">
					<button class="px-3 py-1.5 text-sm border border-transparent text-muted-foreground hover:bg-muted rounded-md disabled:opacity-50 transition-colors font-medium">Prev</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-primary bg-[#2563EB] text-white rounded-md font-semibold">1</button>
					<button class="px-3 py-1.5 text-sm border border-transparent text-foreground hover:bg-muted rounded-md transition-colors font-medium text-muted-foreground hover:text-foreground">Next</button>
				</div>
			</div>
		</div>
	{/await}
</div>

<!-- Add Log Modal -->
<EntryModal 
	isOpen={isEntryModalOpen} 
	title="Add Activity Log" 
	type="log"
	onClose={() => isEntryModalOpen = false} 
	onSave={handleSaveLog} 
/>
