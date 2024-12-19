<script>
	import { Tab, TabGroup } from '@skeletonlabs/skeleton';
	let tabSet = 0
	let innerTab = 0
</script>
<TabGroup>
	<Tab bind:group={tabSet} name="Resources" value={0}>
		<svelte:fragment slot="lead">Resources</svelte:fragment>
	</Tab>
	<Tab bind:group={tabSet} name="Teams" value={1}>Teams</Tab>
	<Tab bind:group={tabSet} name="Sorties" value={2}>Sorties</Tab>

	<!-- Tab Panels --->
	<svelte:fragment slot="panel">
		{#if tabSet === 0}
			<TabGroup>
				<Tab bind:group={innerTab} name="In Progress" value={0}>In Progress</Tab>
				<Tab bind:group={innerTab} name="Unassigned" value={1}>Unassigned</Tab>
				<Tab bind:group={innerTab} name="Completed" value={2}>Completed</Tab>
			</TabGroup>
		{:else if tabSet === 1}
			<TabGroup>
				<Tab bind:group={innerTab} name="In Progress" value={0}>In Progress</Tab>
				<Tab bind:group={innerTab} name="Unassigned" value={1}>Unassigned</Tab>
				<Tab bind:group={innerTab} name="Completed" value={2}>Completed</Tab>
			</TabGroup>
		{:else if tabSet === 2}
			<TabGroup>
				<Tab bind:group={innerTab} name="In Progress" value={0}>In Progress</Tab>
				<Tab bind:group={innerTab} name="Unassigned" value={1}>Unassigned</Tab>
				<Tab bind:group={innerTab} name="Completed" value={2}>Completed</Tab>
			</TabGroup>
		{/if}
	</svelte:fragment>
</TabGroup>