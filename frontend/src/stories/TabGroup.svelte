<script>
	import { TabGroup, Tab } from '@skeletonlabs/skeleton';
	import { fade } from 'svelte/transition';

	export let tabs = [];  // Pass an array of tabs dynamically

	let activeTab = 0;  // Track active tab index
	let expandedTabs = [];  // Track expanded adjacent tabs

	// Expand adjacent tabs
	function expandAdjacentTabs(index) {
		const leftTab = tabs[index - 1]?.value;
		const rightTab = tabs[index + 1]?.value;
		expandedTabs = [leftTab, rightTab].filter(Boolean);
	}

	// Collapse adjacent tabs
	function collapseAdjacentTabs() {
		expandedTabs = [];
	}
</script>

<!-- TabGroup using SkeletonTabs -->
<TabGroup bind:activeTab>
	{#each tabs as { name, value, content }, i}
		<Tab
			bind:activeTab={activeTab}
			value={value}
			on:click={() => {
        activeTab = value;
        expandAdjacentTabs(i);  // Expand adjacent tabs when one tab is clicked
      }}
		>
			{name}
		</Tab>
	{/each}

	<!-- Show expanded tabs if any -->
	{#each expandedTabs as tab}
		<div class="mt-4" transition:fade>
			<TabGroup bind:activeTab>
				<Tab value={tab}>{tabs[tab]?.name} Content</Tab>
			</TabGroup>
		</div>
	{/each}

	<!-- Back button to collapse expanded tabs -->
	{#if expandedTabs.length > 0}
		<button class="mt-4" on:click={collapseAdjacentTabs}>Back</button>
	{/if}
</TabGroup>
