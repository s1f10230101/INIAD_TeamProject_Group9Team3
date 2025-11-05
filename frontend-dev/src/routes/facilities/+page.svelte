<script lang="ts">
  import Facility from "./Facility.svelte";

  let { data } = $props();
  const { facilitiesWithRatings } = data;
</script>

<div class="p-3 flex flex-col bg-white/90 rounded-2xl shadow-2xl mt-2 w-full">
  {#await facilitiesWithRatings}
    <p>loading</p>
  {:then facilities}
    {#each facilities.filter((e) => e !== undefined) as facility}
      <Facility {...facility} />
    {/each}
  {:catch err}
    {err}
  {/await}
</div>
