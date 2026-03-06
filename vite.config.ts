import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  clearScreen: false,
  build: {
    target: "esnext",
  },
  worker: {
    format: "es",
  },
  plugins: [sveltekit()],
});
