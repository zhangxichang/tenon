<script lang="ts">
  import type { Task } from "~/lib/types";
  import type { ViewportStore } from "~/stores/viewport";

  const props: { viewport_store: ViewportStore } = $props();
  let selected_task = $state<Task>();
  let create_task_dialog = $state<HTMLDialogElement>();
</script>

<dialog
  bind:this={create_task_dialog}
  onclick={() => create_task_dialog?.close()}
></dialog>

<div class="flex items-center border-b p-4 justify-between">
  <span class="select-none">任务看板</span>
  <button class="border px-1" onclick={() => create_task_dialog?.showModal()}
    >创建任务</button
  >
</div>
<div class="min-h-0 flex-1 relative flex">
  <div
    class="flex-1 flex flex-col"
    role="button"
    tabindex="0"
    onclick={() => (selected_task = undefined)}
    onkeydown={(e) => e.key === "Enter" && (selected_task = undefined)}
  >
    <div class="overflow-y-auto grid grid-cols-4 gap-4 p-4">
      {#await props.viewport_store.rpcc.get_all_task() then tasks}
        {#each tasks as task (task.id)}
          <div class="flex justify-center">
            <div
              class="flex flex-col border p-2"
              role="button"
              tabindex="0"
              onclick={(e) => {
                e.stopPropagation();
                selected_task = task;
              }}
              onkeydown={(e) => {
                if (e.key === "Enter") {
                  e.stopPropagation();
                  selected_task = task;
                }
              }}
            >
              <span class="text-xs font-bold select-none">#{task.id}</span>
              <span class="font-bold">{task.config.name}</span>
              <span class="text-sm">{task.config.description}</span>
            </div>
          </div>
        {/each}
      {/await}
    </div>
  </div>
  {#if selected_task}
    <div
      class="absolute top-0 right-0 h-full w-80 bg-white border-l flex flex-col"
    >
      <div class="flex-1">
        <div class="flex items-center justify-between border-b p-2">
          <span class="select-none">任务详情</span>
          <button class="border px-1">蓝图</button>
        </div>
        <pre>{selected_task.config.script}</pre>
      </div>
      <div class="border-t flex">
        <button
          class="border-r flex-1"
          onclick={() => (selected_task = undefined)}>关闭</button
        >
        <button class="flex-1">应用</button>
      </div>
    </div>
  {/if}
</div>
