<script lang="ts">
  import { page } from "$app/state";
</script>

<div class="flex-1 flex justify-center items-center">
  <span class="font-bold text-red-400">
    错误: {page.error?.message}
  </span>
</div>
