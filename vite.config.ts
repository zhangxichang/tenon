import { svelte } from "@sveltejs/vite-plugin-svelte";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";

export default defineConfig({
	clearScreen: false,
	resolve: {
		alias: {
			"~": "/src",
		},
	},
	build: {
		target: "esnext",
	},
	worker: {
		format: "es",
	},
	plugins: [tailwindcss(), svelte()],
	server: {
		proxy: {
			"/ws": {
				target: "ws://localhost:10270",
				ws: true,
			},
		},
	},
});
