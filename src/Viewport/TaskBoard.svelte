<script lang="ts">
  import type { ViewportStore } from "~/stores/viewport";

  const props: { viewport_store: ViewportStore } = $props();
  (async () => {
    props.viewport_store.websocket.send("你好，世界");
  })();
</script>

<div></div>
