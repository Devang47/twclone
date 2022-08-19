import { sveltekit } from '@sveltejs/kit/vite';
import { resolve } from 'path'

/** @type {import('vite').UserConfig} */
const config = {
	plugins: [sveltekit()],
	resolve: {
		alias: {
			$components: resolve('./src/lib/components'),
			$icons: resolve('./src/lib/icons'),
			$styles: resolve('./src/styles'),
			$store: resolve('./src/store'),
			$utils: resolve('./src/utils')
		}
	}
};

export default config;
