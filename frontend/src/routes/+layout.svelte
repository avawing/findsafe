<!--This is the Home Page-->
<script>
	import Navbar from '../stories/Navbar.svelte';

	import '../app.css';
	let { children } = $props();

</script>
<main>
	<Navbar />
</main>
	<!-- The rest of your content here -->
{@render children()}
