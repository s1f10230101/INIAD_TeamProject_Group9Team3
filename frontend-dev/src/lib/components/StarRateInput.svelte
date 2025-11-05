<script lang="ts">
    let { star = $bindable(1), name }: { star: number, name: string } = $props();
</script>

<div class="flex text-4xl justify-center py-2">
  {#each { length: 5 }, i}
    {#if i < star}
    <input
      class="hidden"
      type="radio"
      id={i.toString()}
      {name}
      value={i + 1}
      bind:group={star}
    /><label class="text-yellow-300" for={i.toString()}>★</label>
    {:else}
    <input
      class="hidden"
      type="radio"
      {name}
      id={i.toString()}
      value={i + 1}
      bind:group={star}
    /><label for={i.toString()}>★</label>
    {/if}
  {/each}
  <span class="pl-3">
    星:{star}
  </span>
</div>
