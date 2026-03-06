import { defineConfig } from "eslint/config";
import eslint from "@eslint/js";
import ts from "typescript-eslint";
import svelte from "eslint-plugin-svelte";

export default defineConfig({
  files: ["src/**/*.{ts,svelte}"],
  languageOptions: {
    parserOptions: {
      projectService: true,
      parser: ts.parser,
      extraFileExtensions: [".svelte"],
    },
  },
  extends: [
    eslint.configs.recommended,
    ts.configs.recommendedTypeChecked,
    svelte.configs.recommended,
  ],
  rules: {
    "@typescript-eslint/strict-boolean-expressions": "error",
    "@typescript-eslint/no-misused-promises": "off",
    "no-unassigned-vars": "off",
  },
});
