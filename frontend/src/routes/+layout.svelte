<script>
	import { AppBar } from '@skeletonlabs/skeleton';
	import Navbar from '../stories/Navbar.svelte';

	import '../app.css';
	let { children } = $props();
	let valueSingle = 'books';

</script>
<main>


	<Navbar />
</main>
	<!-- The rest of your content here -->
{@render children()}
