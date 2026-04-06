<script lang="ts">
import "~/style.css";
import { ViewportStore } from "~/stores/viewport";
import TaskBoard from "./Viewport/TaskBoard.svelte";

let store = $state<ViewportStore>();
const view = $state<"TaskBoard">("TaskBoard");
$effect(() => {
	let mounted = true;
	(async () => {
		const instance = await ViewportStore.new();
		if (mounted) {
			store = instance;
		} else {
			instance.cleanup();
		}
	})();
	return () => {
		mounted = false;
		store?.cleanup();
	};
});
</script>

<div class="absolute w-dvw h-dvh flex flex-col">
  {#if store}
    {#if view === "TaskBoard"}
      <TaskBoard viewport_store={store} />
    {/if}
  {/if}
</div>
