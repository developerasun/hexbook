import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const VITE_DEFAULT_PORT = '5173';
export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	// @dev development port only
	server: {
		port: parseInt(process.env.PORT ?? VITE_DEFAULT_PORT)
	}
});
