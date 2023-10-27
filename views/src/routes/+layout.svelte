<script lang="ts">
        import {
                
                AppShell,
                AppBar,
                AppRail,
                AppRailAnchor,
                Avatar
                
        } from '@skeletonlabs/skeleton';
	import '../app.postcss';
	import { page } from '$app/stores';
</script>

<AppShell>
	<svelte:fragment slot="header">
	        
<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
	<svelte:fragment slot="lead">(icon)</svelte:fragment>
	(title)
	<svelte:fragment slot="trail">
	        
<Avatar
	border="border-4 border-surface-300-600-token hover:!border-primary-500"
	cursor="cursor-pointer"
/>

	</svelte:fragment>
</AppBar>

	</svelte:fragment>
	<svelte:fragment slot="sidebarLeft">
	        
<AppRail>
	<AppRailAnchor href="/" selected={$page.url.pathname === '/'}>Home</AppRailAnchor>
	<AppRailAnchor href="/post" selected={$page.url.pathname === '/post'}>Post</AppRailAnchor>
</AppRail>

	</svelte:fragment>
	<!-- (sidebarRight) -->
	<!-- (pageHeader) -->
	<!-- Router Slot -->
	<div style="margin: 50px">
	     	<slot />   
	</div>
	
	<!-- ---- / ---- -->
	<svelte:fragment slot="pageFooter">Page Footer</svelte:fragment>
	<!-- (footer) -->
</AppShell>
