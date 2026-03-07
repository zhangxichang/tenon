<script lang="ts">
  type Task = {
    name: string;
    description: string;
  };
  const tasks = $state(new Map<string, Task>());
  for (let index = 0; index < 30; index++) {
    tasks.set(index.toString(), { name: "任务名称", description: "任务描述" });
  }
</script>

<div class="grid grid-cols-4 overflow-y-auto">
  {#each tasks as task (task[0])}
    <div class="flex flex-col">
      <span>{task[0]}</span>
      <span>{task[1].name}</span>
      <span>{task[1].description}</span>
    </div>
  {/each}
</div>
