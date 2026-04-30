<script lang="ts">
	import { X } from 'lucide-svelte';
	import type { Member, ActivityLog, EntryFormData } from '$lib/types';

	let {
		isOpen = false,
		title = 'Add Member',
		type = 'member',
		initialData = null,
		onClose,
		onSave
	} = $props<{
		isOpen: boolean;
		title: string;
		type?: 'member' | 'log';
		initialData?: Member | ActivityLog | null;
		onClose: () => void;
		onSave: (data: EntryFormData) => void;
	}>();

	let formData = $state<EntryFormData>(
		initialData || {
			name: '',
			plate: '',
			rfid: '',
			uhf: '',
			status: 'IN',
			description: ''
		}
	);

	$effect(() => {
		if (isOpen && initialData) {
			formData = { ...initialData };
		} else if (isOpen && !initialData) {
			formData = { name: '', plate: '', rfid: '', uhf: '', status: 'IN', description: '' };
		}
	});

	function handleSubmit(e: Event) {
		e.preventDefault();
		onSave(formData);
	}
</script>

{#if isOpen}
	<div class="fixed inset-0 z-50 flex items-center justify-center">
		<div class="absolute inset-0 bg-black/50 backdrop-blur-sm transition-opacity" onclick={onClose}></div>
		
		<!-- Modal Content -->
		<div class="relative z-50 w-full max-w-md rounded-xl bg-card border border-border shadow-lg overflow-hidden flex flex-col">
			<div class="flex items-center justify-between p-4 border-b border-border bg-muted/20">
				<h2 class="text-lg font-semibold text-foreground tracking-tight">{title}</h2>
				<button class="p-1 rounded-md text-muted-foreground hover:bg-muted transition-colors cursor-pointer" onclick={onClose}>
					<X class="w-5 h-5" />
				</button>
			</div>
			
			<form onsubmit={handleSubmit} class="flex flex-col flex-1">
				<div class="p-4 space-y-4 flex-1 overflow-y-auto max-h-[70vh]">
					<div class="space-y-1.5">
						<label class="text-sm font-medium text-foreground">Name</label>
						<input type="text" bind:value={formData.name} required class="w-full px-3 py-2 bg-secondary/10 border border-transparent rounded-md text-sm focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors" />
					</div>
					
					<div class="space-y-1.5">
						<label class="text-sm font-medium text-foreground">License Plate</label>
						<input type="text" bind:value={formData.plate} required class="w-full px-3 py-2 bg-secondary/10 border border-transparent rounded-md text-sm font-mono focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors" />
					</div>
					
					<div class="space-y-1.5">
						<label class="text-sm font-medium text-foreground">RFID ID</label>
						<input type="text" bind:value={formData.rfid} required class="w-full px-3 py-2 bg-secondary/10 border border-transparent rounded-md text-sm focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors" />
					</div>
					
					<div class="space-y-1.5">
						<label class="text-sm font-medium text-foreground">UHF ID</label>
						<input type="text" bind:value={formData.uhf} required class="w-full px-3 py-2 bg-secondary/10 border border-transparent rounded-md text-sm focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors" />
					</div>
					
					{#if type === 'log'}
						<div class="space-y-1.5">
							<label class="text-sm font-medium text-foreground">Status</label>
							<select bind:value={formData.status} class="w-full px-3 py-2 bg-secondary/10 border border-transparent rounded-md text-sm focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors">
								<option value="IN">IN</option>
								<option value="OUT">OUT</option>
							</select>
						</div>
						
						<div class="space-y-1.5">
							<label class="text-sm font-medium text-foreground">Description</label>
							<textarea bind:value={formData.description} class="w-full px-3 py-2 bg-secondary/10 border border-transparent rounded-md text-sm focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus:bg-background focus:border-border transition-colors min-h-[80px]"></textarea>
						</div>
					{/if}
				</div>
				
				<div class="p-4 border-t border-border bg-muted/20 flex justify-end gap-3">
					<button type="button" class="px-4 py-2 rounded-md text-sm font-medium border border-border bg-background hover:bg-muted text-foreground transition-colors cursor-pointer" onclick={onClose}>
						Cancel
					</button>
					<button type="submit" class="px-4 py-2 rounded-md text-sm font-medium bg-[#2563EB] hover:bg-blue-700 text-white transition-colors shadow-sm cursor-pointer">
						Save
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
