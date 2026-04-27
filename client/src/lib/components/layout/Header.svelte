<script lang="ts">
	import { Bell, Calendar, ChevronDown, Clock, Search, Menu } from 'lucide-svelte';

	let { sidebarOpen = $bindable(), sidebarCollapsed = $bindable() } = $props();

	let searchQuery = $state('');
	
	// Format current date nicely
	const nowDate = new Date('2023-10-24T10:42:00');
	const formattedDate = new Intl.DateTimeFormat('en-US', {
		month: 'short',
		day: 'numeric',
		hour: 'numeric',
		minute: 'numeric',
		hour12: true
	}).format(nowDate);
</script>

<header class="flex h-16 w-full shrink-0 items-center justify-between border-b border-border bg-background px-4 md:px-6">
	<!-- Mobile Menu Toggle -->
	<button class="inline-flex md:hidden mr-4 p-2 text-muted-foreground hover:text-foreground hover:bg-muted rounded-md" onclick={() => sidebarOpen = true}>
		<Menu class="h-5 w-5" />
	</button>
	<!-- Tablet/Desktop Toggle -->
	<button class="hidden md:inline-flex mr-4 p-2 text-muted-foreground hover:text-foreground hover:bg-muted rounded-md" onclick={() => sidebarCollapsed = !sidebarCollapsed}>
		<Menu class="h-5 w-5" />
	</button>

	<div class="flex flex-1 items-center gap-4 min-w-0">
		<div class="relative w-full max-w-md hidden sm:block">
			<Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
			<input
				type="search"
				bind:value={searchQuery}
				placeholder="Search plates, members..."
				class="h-9 w-full rounded-md border-0 bg-secondary/10 pl-9 pr-4 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50"
			/>
		</div>
		<button class="sm:hidden p-2 text-muted-foreground hover:bg-muted rounded text-foreground">
			<Search class="h-5 w-5" />
		</button>
	</div>

	<div class="flex items-center gap-2 md:gap-6 text-sm text-muted-foreground ml-auto pl-4">
		<div class="hidden lg:flex items-center gap-2">
			<Clock class="h-4 w-4" />
			<span class="truncate">{formattedDate}</span>
		</div>

		<div class="flex items-center gap-2 md:gap-4 md:border-r border-border md:pr-6 pr-2">
			<button class="relative text-muted-foreground hover:text-foreground p-1 md:p-0">
				<Bell class="h-5 w-5" />
				<span class="absolute right-0 top-0 h-2 w-2 rounded-full bg-primary ring-2 ring-background"></span>
			</button>
			<button class="hidden sm:block text-muted-foreground hover:text-foreground">
				<Calendar class="h-5 w-5" />
			</button>
		</div>

		<button class="flex items-center gap-2 font-medium text-foreground p-1 shrink-0">
			<img
				src="https://ui-avatars.com/api/?name=Operator+Alpha&background=0D8ABC&color=fff&rounded=true"
				alt="Operator Alpha"
				class="h-8 w-8 rounded-full"
			/>
			<span class="hidden md:inline-block truncate">AdminLPR</span>
			<ChevronDown class="hidden md:inline-block h-4 w-4 text-muted-foreground" />
		</button>
	</div>
</header>
