<script>
	import { AppRailTile } from '@skeletonlabs/skeleton';

	export let label = ''; // The label to display (e.g., 'Home', 'Search')
	export let title = ''; // The title to change to when clicked
	export let href = '#'; // The href to navigate to when clicked
	export let onClick = () => {}; // Optional custom click handler

	// Optional: You can pass a function to be called when the item is clicked (default is an empty function)
	function handleClick() {
		onClick(title); // Call the onClick function passed as prop
	}
</script>

<AppRailTile
	on:click={handleClick}
	href={href}
	class="text-xl hover:bg-primary-hover-token active:bg-primary-active-token p-3 rounded-md"
>
	{label}
</AppRailTile>