import { Server } from '@sveltejs/kit';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const devProxyServer = "http://localhost:1337/";

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		host: "0.0.0.0",
		port: 9999,
		proxy: {
			"^/api": {
				target: devProxyServer,
				changeOrigin: true,
			},
		},
	},
});
