<script lang="ts">
	import { HelpCircle, LayoutGrid, LogOut, Settings, History, Users, X } from 'lucide-svelte';
	import { page } from '$app/state';

	let { isOpen = $bindable(false), isCollapsed = $bindable(false) } = $props();
	let mobileAwareCollapse = $derived(isCollapsed && !isOpen);

	const navItems = [
		{ name: 'Dashboard', icon: LayoutGrid, href: '/' },
		{ name: 'Members', icon: Users, href: '/members' },
		{ name: 'Activity Logs', icon: History, href: '/logs' },
	];
</script>

{#if isOpen}
	<div class="fixed inset-0 z-40 bg-black/50 md:hidden" onclick={() => isOpen = false} aria-hidden="true" role="button" tabindex="-1"></div>
{/if}

<aside class="fixed inset-y-0 left-0 z-50 flex flex-col border-r border-border bg-sidebar transition-all duration-300 md:static md:translate-x-0 h-screen overflow-y-auto overflow-x-hidden {isOpen ? 'translate-x-0' : '-translate-x-full'} {mobileAwareCollapse ? 'w-20' : 'w-64'}">
	<div class="flex h-16 shrink-0 items-center justify-between px-4 {mobileAwareCollapse ? 'justify-center' : ''}">
		<div class="flex items-center gap-3 {mobileAwareCollapse ? 'justify-center' : ''}">
			<div class="flex h-8 w-8 shrink-0 items-center justify-center rounded bg-primary text-primary-foreground font-bold">
				L
			</div>
			{#if !mobileAwareCollapse}
			<div class="flex flex-col">
				<span class="text-lg font-bold leading-none text-foreground truncate">LPRPark</span>
				<span class="text-[10px] font-medium tracking-wider text-muted-foreground uppercase mt-1 truncate">IOT & LPR Infrastructure</span>
			</div>
			{/if}
		</div>
		<button class="md:hidden p-2 -mr-2 text-muted-foreground hover:bg-muted rounded-md" onclick={() => isOpen = false}>
			<X class="h-5 w-5" />
		</button>
	</div>

	<nav class="flex-1 space-y-1 px-4 py-6">
		{#each navItems as item}
			{@const isActive = page.url.pathname === item.href}
			<a
				href={item.href}
				title={mobileAwareCollapse ? item.name : undefined}
				class="group flex items-center gap-3 rounded-md px-3 py-2 text-sm font-medium transition-colors {mobileAwareCollapse ? 'justify-center px-0' : ''} {isActive
					? 'bg-blue-50/50 text-primary relative after:absolute after:left-0 after:top-1/2 after:h-8 after:w-1 after:-translate-y-1/2 after:rounded-r-md after:bg-primary'
					: 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
			>
				<item.icon class="h-5 w-5 shrink-0 {isActive ? 'text-primary' : 'text-neutral'}" />
				{#if !mobileAwareCollapse}
				<span class="truncate">{item.name}</span>
				{/if}
			</a>
		{/each}
	</nav>

	<div class="p-4 space-y-2 shrink-0">
		<a
			href="/support"
			title={mobileAwareCollapse ? 'Support Center' : undefined}
			class="flex items-center {mobileAwareCollapse ? 'justify-center' : 'justify-center'} gap-2 rounded-md bg-blue-50/50 px-3 py-2.5 text-sm font-medium text-primary transition-colors hover:bg-blue-100/50"
		>
			<HelpCircle class="h-4 w-4 shrink-0" />
			{#if !mobileAwareCollapse}
			<span class="truncate">Support Center</span>
			{/if}
		</a>
		<button
			title={mobileAwareCollapse ? 'Sign Out' : undefined}
			class="cursor-pointer flex w-full items-center {mobileAwareCollapse ? 'justify-center px-0' : 'justify-center px-3'} gap-3 rounded-md py-2 text-sm font-medium text-muted-foreground transition-colors hover:bg-muted hover:text-foreground"
		>
			<LogOut class="h-5 w-5 text-neutral shrink-0" />
			{#if !mobileAwareCollapse}
			<span class="truncate">Sign Out</span>
			{/if}
		</button>
	</div>
</aside>
