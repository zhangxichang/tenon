<script lang="ts">
  import type { Task } from "~/lib/types";
  import type { ViewportStore } from "~/stores/viewport";

  const props: { viewport_store: ViewportStore } = $props();
  let selected_task = $state<Task>();
</script>

<div class="flex-1 flex flex-col p-2 gap-2">
  <div class="flex items-center justify-between border p-4">
    <span>任务看板</span>
    <input class="border" placeholder="搜索任务" />
  </div>
  <div class="min-h-0 overflow-y-auto grid grid-cols-4 gap-6 gap-x-32">
    {#await props.viewport_store.rpcc.get_all_task() then tasks}
      {#each tasks as task (task.id)}
        <button
          class="flex flex-col border p-2 items-start"
          onclick={() => (selected_task = task)}
        >
          <span class="text-xs font-bold">#{task.id}</span>
          <span class="font-bold">{task.config.name}</span>
          <span class="text-sm">{task.config.description}</span>
        </button>
      {/each}
    {/await}
  </div>
  {#if selected_task}
    <div class="absolute">{selected_task.config.script}</div>
  {/if}
</div>
