import type { Config } from 'tailwindcss';
import {skeleton} from "@skeletonlabs/tw-plugin";

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {}
	},

	plugins: [
		skeleton({
			themes: { preset: [ "vintage" ] }
		})
	]
} satisfies Config;