<script lang="ts">
  /**
   * @prop {star}
   */
  // star is int 1~5
  let { star }: { star: number } = $props();
  let starSt = $derived(Math.round(star - 1));
</script>

<span class="text-4xl flex justify-center">
  <!-- 10回繰り返す, Svelte 5最新構文 -->
  {#each { length: 5 }, i}
    {#if i <= starSt}
      <span class="text-yellow-500">★</span>
    {:else}
      <span class="text-gray-600">★</span>
    {/if}
  {/each}
  <span class="text-orange-700 font-bold">{star.toFixed(2)}</span>
</span>
