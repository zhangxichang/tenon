<script lang="ts">
  import type { ViewportStore } from "~/stores/viewport";

  type Task = {
    id: string;
    name: string;
    description: string;
  };

  const props: { viewport_store: ViewportStore } = $props();
  const tasks: Task[] = [];
  (async () => {
    await props.viewport_store.rpcc.test();
  })();
</script>

<div class="flex items-center justify-between p-4">
  <span class="font-bold text-3xl select-none">任务看板</span>
  <input class="border rounded" />
</div>
<div class="min-h-0 overflow-y-auto">
  {#each tasks as task (task.id)}
    <button>{task.name}</button>
  {/each}
</div>
