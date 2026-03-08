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
  plugins: [tailwindcss(), svelte()],
});
