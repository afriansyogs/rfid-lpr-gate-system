<script lang="ts">
	import { AlertTriangle } from 'lucide-svelte';

	let {
		isOpen = false,
		title = 'Confirm Action',
		message = 'Are you sure you want to proceed?',
		onClose,
		onConfirm
	} = $props<{
		isOpen: boolean;
		title?: string;
		message?: string;
		onClose: () => void;
		onConfirm: () => void;
	}>();
</script>

{#if isOpen}
	<div class="fixed inset-0 z-50 flex items-center justify-center">
		<div class="absolute inset-0 bg-black/50 backdrop-blur-sm transition-opacity" onclick={onClose}></div>
		
		<!-- Modal Content -->
		<div class="relative z-50 w-full max-w-sm rounded-xl bg-card border border-border shadow-lg overflow-hidden flex flex-col">
			<div class="p-6 flex flex-col items-center text-center space-y-4">
				<div class="w-12 h-12 rounded-full bg-[#BC4800]/10 flex items-center justify-center text-[#BC4800] mb-2">
					<AlertTriangle class="w-6 h-6" />
				</div>
				<h2 class="text-xl font-bold text-foreground tracking-tight">{title}</h2>
				<p class="text-sm text-muted-foreground">{message}</p>
			</div>
			
			<div class="p-4 border-t border-border bg-muted/20 flex justify-end gap-3">
				<button type="button" class="flex-1 px-4 py-2 rounded-md text-sm font-medium border border-border bg-background hover:bg-muted text-foreground transition-colors cursor-pointer" onclick={onClose}>
					Cancel
				</button>
				<button type="button" class="flex-1 px-4 py-2 rounded-md text-sm font-medium bg-[#BC4800] hover:bg-orange-700 text-white transition-colors shadow-sm cursor-pointer" onclick={onConfirm}>
					Delete
				</button>
			</div>
		</div>
	</div>
{/if}
