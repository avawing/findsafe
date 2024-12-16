<script>
	import { AppBar, AppRail, AppRailTile, ListBox, ListBoxItem } from '@skeletonlabs/skeleton';
	import '../app.css';
	let { children } = $props();
	let valueSingle = 'books';

</script>
<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
	<svelte:fragment slot="lead">(icon)</svelte:fragment>
	(title)
	<svelte:fragment slot="trail">(actions)</svelte:fragment>
</AppBar>

<AppRail>
		<AppRailTile bind:group={valueSingle} name="medium" value="books">Books</AppRailTile>
		<AppRailTile bind:group={valueSingle} name="medium" value="movies">Movies</AppRailTile>
		<AppRailTile bind:group={valueSingle} name="medium" value="tv">TV</AppRailTile>
</AppRail>

{@render children()}
