import { defineConfig } from 'vite';  // Correct import for Vite's config
import { sveltekit } from '@sveltejs/kit/vite';  // Correct import for SvelteKit plugin

export default defineConfig({
	plugins: [sveltekit()],
});
