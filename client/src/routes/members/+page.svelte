<script lang="ts">
	import { Search, Download, FileUp, Pen, Trash2, Plus, Copy, Check } from 'lucide-svelte';

	const loadMembersData = () => new Promise(resolve => setTimeout(resolve, 500));
	
	let copiedId = $state<string | null>(null);
	const copyText = async (text: string) => {
		await navigator.clipboard.writeText(text);
		copiedId = text;
		setTimeout(() => {
			if (copiedId === text) copiedId = null;
		}, 2000);
	};

	const members = [
		{ name: 'Sarah Jenkins', plate: 'ABC - 1234', rfid: 'RFID-7712', uhf: 'UHF-0988' },
		{ name: 'Marcus Chen', plate: 'XYZ - 9876', rfid: 'RFID-4421', uhf: 'UHF-3310' },
		{ name: 'Elena Rodriguez', plate: 'LMN - 4567', rfid: 'RFID-8890', uhf: 'UHF-1122' },
		{ name: 'David Kim', plate: 'PQR - 3321', rfid: 'RFID-1029', uhf: 'UHF-7754' },
		{ name: 'Aisha Patel', plate: 'STU - 8899', rfid: 'RFID-5566', uhf: 'UHF-9001' }
	];
</script>

<div class="mx-auto max-w-8xl flex-1 flex flex-col pt-2 pb-10">
	<div class="flex items-center justify-between mb-8">
		<h1 class="text-2xl font-semibold text-foreground tracking-tight">Members</h1>
		<button class="cursor-pointer bg-[#2563EB] text-white px-4 py-2.5 rounded-md font-medium text-sm flex items-center gap-2 hover:bg-blue-700 transition-colors shadow-sm focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2">
			<Plus class="w-4 h-4" /> Add Member
		</button>
	</div>

	{#await loadMembersData()}
		<div class="animate-pulse flex flex-col w-full">
			<div class="h-16 bg-muted/40 rounded-t-xl border border-border"></div>
			<div class="h-[400px] bg-muted/30 border-x border-border"></div>
			<div class="h-16 bg-muted/40 rounded-b-xl border border-border"></div>
		</div>
	{:then}
		<div class="flex flex-col w-full">
			<!-- Toolbar -->
			<div class="flex flex-col sm:flex-row justify-between items-start sm:items-center p-4 border border-border rounded-t-xl bg-card gap-4">
				<div class="relative w-full sm:max-w-md">
					<Search class="absolute left-3 top-2.5 h-4 w-4 text-muted-foreground" />
					<input type="text" placeholder="Search members by name, ID or plate..." class="pl-9 pr-4 py-2 bg-secondary/10 border-0 rounded-md w-full text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring placeholder:text-muted-foreground" />
				</div>
				<div class="flex items-center gap-3 w-full sm:w-auto">
					<button class="cursor-pointer border border-border text-foreground hover:bg-muted px-4 py-2 rounded-md font-medium text-sm flex items-center gap-2 flex-1 justify-center sm:flex-none transition-colors">
						<Download class="w-4 h-4 text-muted-foreground" /> Export
					</button>
					<button class="cursor-pointer border border-border bg-blue-800 text-white hover:bg-blue-900 px-4 py-2 rounded-md font-medium text-sm flex items-center gap-2 flex-1 justify-center sm:flex-none transition-colors">
						<FileUp class="w-4 h-4 text-white" /> Import Data
					</button>
				</div>
			</div>

			<!-- Table -->
			<div class="overflow-x-auto border-x border-b border-border bg-card">
				<table class="w-full text-sm text-left whitespace-nowrap">
					<thead class="text-xs text-muted-foreground uppercase border-b border-border bg-card/50">
						<tr>
							<th class="px-6 py-5 font-extrabold tracking-wider">NO.</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">NAME</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">LICENSE PLATE</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">RFID_ID</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">UHF_ID</th>
							<th class="px-6 py-5 font-extrabold tracking-wider text-right">ACTIONS</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						{#each members as member, i}
							<tr class="hover:bg-muted/30 transition-colors group">
								<td class="px-6 py-4 text-muted-foreground font-medium">{String(i + 1).padStart(2, '0')}</td>
								<td class="px-6 py-4 font-semibold text-foreground">{member.name}</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => copyText(member.plate)} title="Copy Plate">
										<span class="px-2.5 py-1.5 rounded bg-gray-300 border border-gray-400 font-medium text-foreground tracking-widest text-[11px] font-mono">{member.plate}</span>
											{#if copiedId === member.plate}
												<Check class="w-3.5 h-3.5 text-green-500" />
											{:else}
												<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
											{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => copyText(member.rfid)} title="Copy RFID">
										<span>{member.rfid}</span>
										{#if copiedId === member.rfid}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => copyText(member.uhf)} title="Copy UHF ID">
										<span>{member.uhf}</span>
										{#if copiedId === member.uhf}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4 flex items-center justify-end gap-2 text-right">
									<button class="p-2 rounded-md bg-blue-200 hover:bg-blue-300 hover:text-blue-500 text-blue-500 transition-colors cursor-pointer" title="Edit">
										<Pen class="w-4 h-4" />
									</button>
									<button class="p-2 rounded-md bg-red-200 hover:bg-red-300 hover:text-red-500 text-red-500 transition-colors cursor-pointer" title="Delete">
										<Trash2 class="w-4 h-4" />
									</button>
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>

			<!-- Pagination -->
			<div class="flex flex-col sm:flex-row items-center justify-between p-4 border border-t-0 border-border rounded-b-xl bg-card gap-4">
				<span class="text-sm text-muted-foreground font-medium">Showing 1 to 5 of 24 members</span>
				<div class="flex items-center space-x-1.5">
					<button class="px-3 py-1.5 text-sm border border-transparent text-muted-foreground hover:bg-muted rounded-md disabled:opacity-50 transition-colors font-medium">Prev</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-primary bg-[#2563EB] text-white rounded-md font-medium">1</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-transparent text-foreground hover:bg-muted hover:border-border rounded-md font-medium transition-colors">2</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-transparent text-foreground hover:bg-muted hover:border-border rounded-md font-medium transition-colors">3</button>
					<button class="px-3 py-1.5 text-sm border border-transparent text-foreground hover:bg-muted rounded-md transition-colors font-medium text-muted-foreground hover:text-foreground">Next</button>
				</div>
			</div>
		</div>
	{/await}
</div>
