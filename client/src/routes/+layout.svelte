<script lang="ts">
  import '../app.css';
  import favicon from '$lib/assets/favicon.svg';
  import * as NavigationMenu from '$lib/components/ui/navigation-menu/index.js';
  import LoginIcon from '@lucide/svelte/icons/log-in';
  import ProfileIcon from '@lucide/svelte/icons/cat';
  import Darkmode from '$lib/components/ui/Darkmode/darkmode.svelte';
  import Spacer from '$lib/components/ui/spacer/spacer.svelte';
  import { Button } from '$lib/components/ui/button';

  let isLogin = false;
  let { children } = $props();

  const JOURNAL_LIST: { href: string; title: string }[] = [
    {
      href: 'weights',
      title: 'Weights'
    },
    {
      href: 'meals',
      title: 'Meals'
    },
    {
      href: 'exercises',
      title: 'Exercises'
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
        <NavigationMenu.Trigger>Journal</NavigationMenu.Trigger>
        <NavigationMenu.Content class="absolute left-0 top-full">
          <ul>
            {#each JOURNAL_LIST as { href, title }, i (i)}
              {@render NavItem({ href: `/journal/${href}`, title })}
            {/each}
          </ul>
        </NavigationMenu.Content>
      </NavigationMenu.Item>
      <NavigationMenu.Item>
        {@render NavItem({ href: '/dashboard', title: 'Dashboard' })}
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
        <Button href={isLogin ? '/profile' : '/login'} variant="outline" size="icon">
          {#if isLogin}
            <ProfileIcon class="h-[1.2rem] w-[1.2rem] " />
          {:else}
            <LoginIcon class="h-[1.2rem] w-[1.2rem] " />
          {/if}
        </Button>
      </NavigationMenu.Item>
    </NavigationMenu.List>
  </div>
</NavigationMenu.Root>
<Spacer my="my-4" />

<div class="m-auto w-3/4">
  {@render children?.()}
</div>
