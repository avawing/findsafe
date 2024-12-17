<script>
	let isOpen = false;
	import { AppBar, AppRail, AppRailAnchor, AppRailTile } from '@skeletonlabs/skeleton';
	import 'font-awesome/css/font-awesome.min.css';
	let title = 'Home';

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
	class={`fixed inset-0 bg-surface-100-800-token transform ${isOpen ? 'translate-x-0 opacity-100' : '-translate-x-full opacity-0'} transition-all duration-300 ease-in-out z-40`}
>
	<!-- AppTile items inside AppRail -->
	<AppRail class="w-40 h-full bg-surface-100-800-token p-5 flex flex-col gap-0 space-y-1">
		<AppRailAnchor on:click={toggleNavbar} > Back </AppRailAnchor>
		<AppRailTile
			on:click={() => changeTitle('Home') }
			href="#"
			class="text-xl  hover:bg-primary-hover-token active:bg-primary-active-token p-3 rounded-md ">
			Home

		</AppRailTile>
		<AppRailTile
			on:click={() => changeTitle('Search')}
			href="#"
			class="text-xl  hover:bg-primary-hover-token active:bg-primary-active-token p-3 rounded-md ">
			Search
		</AppRailTile>
		<AppRailTile
			on:click={() => changeTitle('Sortie')}
			href="#"
			class="text-xl  hover:bg-primary-hover-token active:bg-primary-active-token p-3 rounded-md ">
			Sortie
		</AppRailTile>
		<AppRailTile
			on:click={() => changeTitle('Resources')}
			href="#"
			class="text-xl  hover:bg-primary-hover-token active:bg-primary-active-token p-3 rounded-md ">
			Resources
		</AppRailTile>
		<AppRailTile
			on:click={() => changeTitle('Organization')}
			href="#"
			class="text-xl  hover:bg-primary-hover-token active:bg-primary-active-token p-3 rounded-md ">
			Organization
		</AppRailTile>
	</AppRail>
</div>