<script lang="ts">
  type Task = {
    id: string;
    name: string;
    description: string;
  };
  const tasks: Task[] = [];
</script>

<div class="flex items-center justify-between p-4">
  <span class="font-bold text-3xl select-none">任务看板</span>
  <input class="border rounded" />
</div>
<div class="min-h-0 overflow-y-auto">
  {#each tasks as task (task.id)}
    <div>{task.name}</div>
  {/each}
</div>
