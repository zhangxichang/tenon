<script lang="ts">
  import "~/style.css";
  import icon from "~/icon.svg";

  const { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={icon} />
  <title>Tenon</title>
</svelte:head>

{@render children()}
