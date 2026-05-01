<script lang="ts">
	import { X, User, Car, CreditCard, Radio, Activity, FileText, Image as ImageIcon, Clock, UploadCloud } from 'lucide-svelte';
	import { fade, scale } from 'svelte/transition';
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
			description: 'success',
			timestamp: '',
			image: ''
		}
	);

	$effect(() => {
		if (isOpen && initialData) {
			formData = { ...initialData };
		} else if (isOpen && !initialData) {
			formData = { 
				name: '', 
				plate: '', 
				rfid: '', 
				uhf: '', 
				status: 'IN', 
				description: 'success',
				timestamp: new Date().toISOString().slice(0, 16),
				image: ''
			};
		}
	});

	function handleSubmit(e: Event) {
		e.preventDefault();
		onSave(formData);
	}
</script>

{#if isOpen}
	<div class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-0" transition:fade={{ duration: 200 }}>
		<div class="absolute inset-0 bg-background/20 backdrop-blur-md transition-opacity" onclick={onClose}></div>
		
		<!-- Modal Content -->
		<div 
			class="relative z-50 w-full max-w-lg rounded-2xl bg-card border border-border/50 shadow-2xl overflow-hidden flex flex-col"
			transition:scale={{ duration: 300, start: 0.95, opacity: 0 }}
		>
			<div class="flex items-center justify-between p-5 border-b border-border/50 bg-muted/30">
				<div class="flex items-center gap-3">
					<div class="p-2 bg-primary/10 rounded-lg">
						{#if type === 'member'}
							<User class="w-5 h-5 text-primary" />
						{:else}
							<Activity class="w-5 h-5 text-primary" />
						{/if}
					</div>
					<h2 class="text-xl font-semibold text-foreground tracking-tight">{title}</h2>
				</div>
				<button 
					class="p-2 rounded-full text-muted-foreground hover:bg-muted hover:text-foreground transition-all cursor-pointer focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring" 
					onclick={onClose}
				>
					<X class="w-5 h-5" />
				</button>
			</div>
			
			<form onsubmit={handleSubmit} class="flex flex-col flex-1 max-h-[85vh]">
				<div class="p-5 space-y-5 flex-1 overflow-y-auto custom-scrollbar">
					{#if type === 'log'}
						<!-- Image Upload Area -->
						<div class="space-y-2">
							<label class="text-sm font-medium text-foreground flex items-center gap-2">
								<ImageIcon class="w-4 h-4 text-muted-foreground" />
								Image - Plate
							</label>
							<div class="border-2 border-dashed border-border rounded-xl p-6 flex flex-col items-center justify-center gap-2 bg-muted/10 hover:bg-muted/30 transition-colors cursor-pointer group">
								<div class="p-3 bg-background rounded-full shadow-sm group-hover:scale-110 transition-transform">
									<UploadCloud class="w-6 h-6 text-muted-foreground" />
								</div>
								<div class="text-center">
									<p class="text-sm font-medium text-foreground">Click to upload or drag and drop</p>
									<p class="text-xs text-muted-foreground mt-1">SVG, PNG, JPG or GIF (max. 800x400px)</p>
								</div>
								<input type="file" class="hidden" accept="image/*" />
							</div>
						</div>
					{/if}

					<div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
						<div class="space-y-2">
							<label class="text-sm font-medium text-foreground flex items-center gap-2">
								<User class="w-4 h-4 text-muted-foreground" />
								Name
							</label>
							<input 
								type="text" 
								bind:value={formData.name} 
								required 
								placeholder="Enter name"
								class="w-full px-4 py-2.5 bg-background border border-input rounded-lg text-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm" 
							/>
						</div>
						<div class="space-y-2">
							<label class="text-sm font-medium text-foreground flex items-center gap-2">
								<Car class="w-4 h-4 text-muted-foreground" />
								License Plate
							</label>
							<input 
								type="text" 
								bind:value={formData.plate} 
								required 
								placeholder="e.g. B 1234 XYZ"
								class="w-full px-4 py-2.5 bg-background border border-input rounded-lg text-sm font-mono uppercase focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm" 
							/>
						</div>
						<div class="space-y-2">
							<label class="text-sm font-medium text-foreground flex items-center gap-2">
								<CreditCard class="w-4 h-4 text-muted-foreground" />
								RFID ID
							</label>
							<input 
								type="text" 
								bind:value={formData.rfid} 
								required 
								placeholder="Scan or enter RFID"
								class="w-full px-4 py-2.5 bg-background border border-input rounded-lg text-sm font-mono focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm" 
							/>
						</div>
						<div class="space-y-2">
							<label class="text-sm font-medium text-foreground flex items-center gap-2">
								<Radio class="w-4 h-4 text-muted-foreground" />
								UHF ID
							</label>
							<input 
								type="text" 
								bind:value={formData.uhf} 
								required 
								placeholder="Scan or enter UHF"
								class="w-full px-4 py-2.5 bg-background border border-input rounded-lg text-sm font-mono focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm" 
							/>
						</div>
					</div>
					
					{#if type === 'log'}
						<div class="space-y-2">
							<label class="text-sm font-medium text-foreground flex items-center gap-2">
								<Clock class="w-4 h-4 text-muted-foreground" />
								Timestamp
							</label>
							<input 
								type="datetime-local" 
								bind:value={formData.timestamp} 
								required
								class="w-full px-4 py-2.5 bg-background border border-input rounded-lg text-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm" 
							/>
						</div>
						<div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
							<div class="space-y-2">
								<label class="text-sm font-medium text-foreground flex items-center gap-2">
									<Activity class="w-4 h-4 text-muted-foreground" />
									Status
								</label>
								<div class="relative">
									<select 
										bind:value={formData.status} 
										class="w-full appearance-none px-4 py-2.5 bg-background border border-input rounded-lg text-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm"
									>
										<option value="IN">IN</option>
										<option value="OUT">OUT</option>
									</select>
									<div class="absolute inset-y-0 right-3 flex items-center pointer-events-none text-muted-foreground">
										<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
									</div>
								</div>
							</div>
							<div class="space-y-2">
								<label class="text-sm font-medium text-foreground flex items-center gap-2">
									<FileText class="w-4 h-4 text-muted-foreground" />
									Description
								</label>
								<div class="relative">
									<select 
										bind:value={formData.description} 
										class="w-full appearance-none px-4 py-2.5 bg-background border border-input rounded-lg text-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:border-transparent transition-all shadow-sm"
									>
										<option value="success">Success</option>
										<option value="RFID not registered">RFID Not Registered</option>
										<option value="plate does not match">Plate Does Not Match</option>
									</select>
									<div class="absolute inset-y-0 right-3 flex items-center pointer-events-none text-muted-foreground">
										<svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path></svg>
									</div>
								</div>
							</div>
						</div>
					{/if}
				</div>
				
				<div class="p-5 border-t border-border/50 bg-muted/30 flex justify-end gap-3 rounded-b-2xl">
					<button 
						type="button" 
						class="px-5 py-2.5 rounded-lg text-sm font-medium border border-input bg-background hover:bg-muted text-foreground transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring cursor-pointer" 
						onclick={onClose}
					>
						Cancel
					</button>
					<button 
						type="submit" 
						class="px-5 py-2.5 rounded-lg text-sm font-medium bg-primary hover:bg-primary/90 text-primary-foreground transition-all shadow-sm focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring cursor-pointer"
					>
						Save {type === 'member' ? 'Member' : 'Log'}
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}

<style>
	.custom-scrollbar::-webkit-scrollbar {
		width: 6px;
	}
	.custom-scrollbar::-webkit-scrollbar-track {
		background: transparent;
	}
	.custom-scrollbar::-webkit-scrollbar-thumb {
		background-color: hsl(var(--border));
		border-radius: 10px;
	}
	.custom-scrollbar:hover::-webkit-scrollbar-thumb {
		background-color: hsl(var(--muted-foreground) / 0.5);
	}
</style>
