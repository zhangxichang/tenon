<script lang="ts">
  import "~/style.css";
  import icon from "~/icon.svg";
  import type { Snippet } from "svelte";

  const props: {
    children: Snippet;
  } = $props();
</script>

<svelte:head>
  <link rel="icon" href={icon} />
  <title>Tenon</title>
</svelte:head>

<div class="absolute w-dvw h-dvh flex flex-col">
  {@render props.children()}
</div>
