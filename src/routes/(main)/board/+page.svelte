<script lang="ts">
  import type { LayoutProps } from "../$types";

  const props: LayoutProps = $props();
  (() => {
    props.data.main_store.websocket.send("你好，世界");
  })();
</script>

<div></div>
