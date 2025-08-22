<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import * as NavigationMenu from '$lib/components/ui/navigation-menu/index.js';
	import LoginIcon from '@lucide/svelte/icons/log-in';
	import Darkmode from '$lib/components/ui/Darkmode/darkmode.svelte';
	import { Button } from '$lib/components/ui/button';

	let { children } = $props();

	const DASHBOARD_LIST: { href: string; title: string }[] = [
		{
			href: '/profile',
			title: 'Profile'
		},
		{
			href: '/weights',
			title: 'Weights'
		},
		{
			href: '/visualization',
			title: 'Visualization'
		}
	];
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{#snippet NavItem({ href, title }: { href: string; title: string })}
	<NavigationMenu.Link {href} class="px-4 py-2 hover:bg-gray-100" rel="noopener noreferrer">
		{title}
	</NavigationMenu.Link>
{/snippet}

<NavigationMenu.Root class="flex flex-col items-center justify-center" viewport={false}>
	<div class="flex w-screen items-center justify-around bg-amber-300 py-4">
		<NavigationMenu.List>
			<NavigationMenu.Item>
				<a href="/">Logo</a>
			</NavigationMenu.Item>
		</NavigationMenu.List>

		<NavigationMenu.List>
			<NavigationMenu.Item>
				<NavigationMenu.Trigger>Dashboard</NavigationMenu.Trigger>
				<NavigationMenu.Content class="absolute left-0 top-full">
					<ul>
						{#each DASHBOARD_LIST as { href, title }, i (i)}
							{@render NavItem({ href, title })}
						{/each}
					</ul>
				</NavigationMenu.Content>
			</NavigationMenu.Item>
			<NavigationMenu.Item>
				{@render NavItem({ href: '/pricing', title: 'Pricing' })}
			</NavigationMenu.Item>
		</NavigationMenu.List>
		<NavigationMenu.List>
			<NavigationMenu.Item>
				<Darkmode />
			</NavigationMenu.Item>
			<NavigationMenu.Item>
				<Button href="/login" variant="outline" size="icon">
					<LoginIcon class="h-[1.2rem] w-[1.2rem] " />
				</Button>
			</NavigationMenu.Item>
		</NavigationMenu.List>
	</div>
</NavigationMenu.Root>
{@render children?.()}
