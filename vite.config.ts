import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  clearScreen: false,
  build: {
    target: "esnext",
  },
  worker: {
    format: "es",
  },
  server: {
    proxy: {
      "/ws": {
        target: "ws://localhost:10270",
        ws: true,
      },
    },
  },
  plugins: [tailwindcss(), svelte()],
  resolve: {
    alias: {
      "~": "/src",
    },
  },
});
