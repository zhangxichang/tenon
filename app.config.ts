import {
  defineConfig,
  type SolidStartInlineConfig,
} from "@solidjs/start/config";
import tailwindcss from "@tailwindcss/vite";

export const config: SolidStartInlineConfig = {
  ssr: false,
  server: {
    preset: "static",
    esbuild: { options: { target: "esnext" } },
  },
  vite: {
    clearScreen: false,
    build: {
      target: "esnext",
    },
    worker: {
      format: "es",
    },
    plugins: [tailwindcss()],
  },
};
export default defineConfig(config);
