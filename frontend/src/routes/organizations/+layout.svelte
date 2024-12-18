<script>
	let tabSet = 0;
	import { Tab, TabGroup } from '@skeletonlabs/skeleton';
</script>

<TabGroup>
	<Tab bind:group={tabSet} name="Members" value={0}>
		<svelte:fragment slot="lead">Members</svelte:fragment>
	</Tab>
	<Tab bind:group={tabSet} name="Resources" value={1}>Resources</Tab>
	<Tab bind:group={tabSet} name="Certifications" value={2}>Certifications</Tab>
	<!-- Tab Panels --->
	<svelte:fragment slot="panel">
		{#if tabSet === 0}
			<main></main>
			(tab panel 1 contents)
		{:else if tabSet === 1}
			(tab panel 2 contents)
		{:else if tabSet === 2}
			(tab panel 3 contents)
		{/if}
	</svelte:fragment>
</TabGroup>

