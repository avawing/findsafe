<script>
	import { AppBar, AppRail, AppRailAnchor} from '@skeletonlabs/skeleton';
	import NavItem from './NavItem.svelte';
	import 'font-awesome/css/font-awesome.min.css';
	let title = 'Home';
	let isOpen = false;

	let navItems = [
		{ label: 'Home', title: 'Home', href: '/' },
		{ label: 'Search', title: 'Search', href: '/searches' },
		{ label: 'Sortie', title: 'Sortie', href: '/sortie' },
		{ label: 'Resources', title: 'Resources', href: '/resources' },
		{ label: 'Organization', title: 'Organization', href: '/organizations' },
	];

  // Dynamic Title based on
	function changeTitle(newTitle) {
		title = newTitle;
		toggleNavbar()
	}
	// Toggle the sidebar visibility
	function toggleNavbar() {
		isOpen = !isOpen;
	}
</script>

<style>
    /* Optional: you can adjust the navbar width here */
</style>

<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
	<!-- Button to toggle navbar -->
	<svelte:fragment slot="lead"><button
		on:click={toggleNavbar}
		class="text-white bg-primary-hover-token rounded-full">
		<i class="fas fa-bars"></i> <!-- Font Awesome Hamburger icon (Optional) -->
	</button></svelte:fragment>
	{title}
	<svelte:fragment slot="trail">(actions)</svelte:fragment>
</AppBar>


<!-- AppRail Sidebar using SkeletonLabs components -->
<div
	class={`fixed inset-0 bg-surface-100-800-token transform ${isOpen ? 'translate-x-0 opacity-100' : '-translate-x-full opacity-0'} transition-all duration-300 ease-in-out z-40`}>

	<AppRail class="w-40 h-full bg-surface-100-800-token p-5 flex flex-col gap-0 space-y-1">
		<AppRailAnchor on:click={toggleNavbar}> Back </AppRailAnchor>

		<!-- Loop over navItems and create NavItem components dynamically -->
		{#each navItems as item (item.label)}
			<a href={item.href}>
			<NavItem
				label={item.label}
				title={item.title}
				onClick={changeTitle}
			/>
			</a>
		{/each}
	</AppRail>
</div>