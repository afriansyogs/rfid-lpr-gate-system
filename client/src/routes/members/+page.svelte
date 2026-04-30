<script lang="ts">
	import { Search, Download, FileUp, Pen, Trash2, Plus, Copy, Check, Loader2 } from 'lucide-svelte';
	import dummyData from '$lib/data/dummy.json';
	import EntryModal from '$lib/components/shared/EntryModal.svelte';
	import ConfirmModal from '$lib/components/shared/ConfirmModal.svelte';
	import { useClipboard } from '$lib/utils/clipboard.svelte';
	import type { Member, EntryFormData } from '$lib/types';

	const loadMembersData = () => new Promise(resolve => setTimeout(resolve, 500));
	const clipboard = useClipboard();

	let members = $state<Member[]>(dummyData.members as Member[]);

	// Modals State
	let isEntryModalOpen = $state(false);
	let isDeleteModalOpen = $state(false);
	let selectedMember = $state<Member | null>(null);
	let entryModalTitle = $state('Add Member');

	// Search State
	let searchQuery = $state('');
	let isSearching = $state(false);
	let searchTimeout: ReturnType<typeof setTimeout>;

	let filteredMembers = $derived(members.filter(m => {
		if (!searchQuery) return true;
		const q = searchQuery.toLowerCase();
		return m.name.toLowerCase().includes(q) ||
						m.plate.toLowerCase().includes(q) ||
						m.rfid.toLowerCase().includes(q) ||
						m.uhf.toLowerCase().includes(q);
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

	function openAddModal() {
		entryModalTitle = 'Add Member';
		selectedMember = null;
		isEntryModalOpen = true;
	}

	function openEditModal(member: Member) {
		entryModalTitle = 'Edit Member';
		selectedMember = member;
		isEntryModalOpen = true;
	}

	function openDeleteModal(member: Member) {
		selectedMember = member;
		isDeleteModalOpen = true;
	}

	function handleSaveMember(data: EntryFormData) {
		if (selectedMember) {
			// Update
			members = members.map(m => m === selectedMember ? { ...m, ...data } as Member : m);
		} else {
			// Add
			members = [data as Member, ...members];
		}
		isEntryModalOpen = false;
	}

	function confirmDelete() {
		if (selectedMember) {
			members = members.filter(m => m !== selectedMember);
		}
		isDeleteModalOpen = false;
	}
</script>

<div class="mx-auto max-w-8xl flex-1 flex flex-col pt-2 pb-10">
	<div class="flex items-center justify-between mb-8">
		<h1 class="text-2xl font-semibold text-foreground tracking-tight">Members</h1>
		<button onclick={openAddModal} class="cursor-pointer bg-[#2563EB] text-white px-4 py-2.5 rounded-md font-medium text-sm flex items-center gap-2 hover:bg-blue-700 transition-colors shadow-sm focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2">
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
					<input 
						type="text" 
						oninput={handleSearchInput}
						placeholder="Search members by name, ID or plate..." 
						class="pl-9 pr-10 py-2 bg-secondary/10 border-0 rounded-md w-full text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring placeholder:text-muted-foreground" 
					/>
					{#if isSearching}
						<Loader2 class="absolute right-3 top-2.5 h-4 w-4 text-muted-foreground animate-spin" />
					{/if}
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
							<th class="px-6 py-5 font-extrabold tracking-wider">RFID ID</th>
							<th class="px-6 py-5 font-extrabold tracking-wider">UHF ID</th>
							<th class="px-6 py-5 font-extrabold tracking-wider text-right">ACTIONS</th>
						</tr>
					</thead>
					<tbody class="divide-y divide-border">
						{#each filteredMembers as member, i}
							<tr class="hover:bg-muted/30 transition-colors group">
								<td class="px-6 py-4 text-muted-foreground font-medium">{String(i + 1).padStart(2, '0')}</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(member.name)} title="Copy Name">
										<span class="font-semibold text-foreground">{member.name}</span>
										{#if clipboard.copiedId === member.name}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(member.plate)} title="Copy Plate">
										<span class="px-2.5 py-1.5 rounded bg-gray-300 border border-gray-400 font-medium text-foreground tracking-widest text-[11px] font-mono">{member.plate}</span>
										{#if clipboard.copiedId === member.plate}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(member.rfid)} title="Copy RFID">
										<span>{member.rfid}</span>
										{#if clipboard.copiedId === member.rfid}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4">
									<button class="group/btn inline-flex w-fit items-center gap-2 text-muted-foreground font-medium tracking-wide hover:text-[#2563EB] transition-colors cursor-pointer focus:outline-none" onclick={() => clipboard.copyText(member.uhf)} title="Copy UHF ID">
										<span>{member.uhf}</span>
										{#if clipboard.copiedId === member.uhf}
											<Check class="w-3.5 h-3.5 text-green-500" />
										{:else}
											<Copy class="w-3.5 h-3.5 opacity-0 group-hover/btn:opacity-100 transition-opacity text-[#2563EB]" />
										{/if}
									</button>
								</td>
								<td class="px-6 py-4 flex items-center justify-end gap-2 text-right">
									<button onclick={() => openEditModal(member)} class="p-2 rounded-md bg-blue-200 hover:bg-blue-300 hover:text-blue-500 text-blue-500 transition-colors cursor-pointer" title="Edit">
										<Pen class="w-4 h-4" />
									</button>
									<button onclick={() => openDeleteModal(member)} class="p-2 rounded-md bg-red-200 hover:bg-red-300 hover:text-red-500 text-red-500 transition-colors cursor-pointer" title="Delete">
										<Trash2 class="w-4 h-4" />
									</button>
								</td>
							</tr>
						{:else}
							<tr>
								<td colspan="6" class="px-6 py-8 text-center text-muted-foreground">
									No members found.
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>

			<!-- Pagination -->
			<div class="flex flex-col sm:flex-row items-center justify-between p-4 border border-t-0 border-border rounded-b-xl bg-card gap-4">
				<span class="text-sm text-muted-foreground font-medium">Showing {filteredMembers.length} members</span>
				<div class="flex items-center space-x-1.5">
					<button class="px-3 py-1.5 text-sm border border-transparent text-muted-foreground hover:bg-muted rounded-md disabled:opacity-50 transition-colors font-medium">Prev</button>
					<button class="w-8 h-8 flex items-center justify-center text-sm border border-primary bg-[#2563EB] text-white rounded-md font-medium">1</button>
					<button class="px-3 py-1.5 text-sm border border-transparent text-foreground hover:bg-muted rounded-md transition-colors font-medium text-muted-foreground hover:text-foreground">Next</button>
				</div>
			</div>
		</div>
	{/await}
</div>

<!-- Modals -->
<EntryModal 
	isOpen={isEntryModalOpen} 
	title={entryModalTitle} 
	type="member"
	initialData={selectedMember}
	onClose={() => isEntryModalOpen = false} 
	onSave={handleSaveMember} 
/>

<ConfirmModal 
	isOpen={isDeleteModalOpen}
	title="Delete Member"
	message="Are you sure you want to delete this member? This action cannot be undone."
	onClose={() => isDeleteModalOpen = false}
	onConfirm={confirmDelete}
/>
