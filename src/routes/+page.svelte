<script lang="ts">
  import TaskBoard from "~/components/TaskBoard.svelte";

  type View = "TaskBoard";
  const view = $state<View>("TaskBoard");
</script>

{#if view === "TaskBoard"}
  <TaskBoard />
{/if}
