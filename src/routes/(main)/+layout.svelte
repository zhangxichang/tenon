<script lang="ts">
  import { onDestroy, type Snippet } from "svelte";
  import type { LayoutProps } from "./$types";

  const props: { children: Snippet } & LayoutProps = $props();
  onDestroy(() => {
    props.data.main_store.cleanup();
  });
</script>

{@render props.children()}
