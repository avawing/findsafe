<script>
	let tabSet = 0;
	import { Tab, TabGroup } from '@skeletonlabs/skeleton';
</script>

<TabGroup>
	<Tab bind:group={tabSet} name="Active Search" value={0}>
		<svelte:fragment slot="lead">Active Search</svelte:fragment>
	</Tab>
	<Tab bind:group={tabSet} name="Active Searches" value={1}>Active Searches</Tab>
	<Tab bind:group={tabSet} name="Past Searches" value={2}>Past Searches</Tab>
	<!-- Tab Panels --->
	<svelte:fragment slot="panel">
		{#if tabSet === 0}
			<main>
				<ul>
					<li><a href="/search/one">one</a></li>
					<li><a href="/search/two">two</a></li>
					<li><a href="/search/three">three</a></li>
				</ul>
			</main>
		{:else if tabSet === 1}
			<ul>
				<li><a href="/search/one">one</a></li>
				<li><a href="/search/two">two</a></li>
				<li><a href="/search/three">three</a></li>
			</ul>
		{:else if tabSet === 2}
			<ul>
				<li><a href="/search/one">one</a></li>
				<li><a href="/search/two">two</a></li>
				<li><a href="/search/three">three</a></li>
			</ul>
		{/if}
	</svelte:fragment>
</TabGroup>



This Page will Host
1. All Active Searches (Organization)
2. All Past Searches (Organization + User)
3. Current Active Search Data (Search Admin)