import eslint from "@eslint/js";
import { defineConfig } from "eslint/config";
import svelte from "eslint-plugin-svelte";
import globals from "globals";
import ts from "typescript-eslint";

export default defineConfig({
	files: ["src/**/*.{ts,svelte}"],
	languageOptions: {
		parserOptions: {
			projectService: true,
			parser: ts.parser,
			extraFileExtensions: [".svelte"],
		},
		globals: { ...globals.browser, ...globals.node },
	},
	extends: [
		eslint.configs.recommended,
		ts.configs.recommended,
		svelte.configs.recommended,
	],
	rules: {
		"@typescript-eslint/strict-boolean-expressions": "error",
		"@typescript-eslint/no-misused-promises": "off",
		"no-unassigned-vars": "off",
	},
});
