<script lang="ts">
import FacilityCard from "./FacilityCard.svelte";
import FacilityCardLoading from "./FacilityCardLoading.svelte";

let { data } = $props();
const { facilitiesWithRatings } = data;
</script>

<div class="p-3 flex flex-col rounded-2xl mt-2 w-full">
  {#await facilitiesWithRatings}
    {#each {length:2} }
     <FacilityCardLoading />
    {/each}
  {:then facilities}
    {#each facilities.filter((e) => e !== undefined) as facility}
      <FacilityCard {...facility} />
    {/each}
  {:catch err}
    {err}
  {/await}
</div>
