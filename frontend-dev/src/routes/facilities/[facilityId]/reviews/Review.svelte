<script lang="ts">
import StarsRate from "$lib/components/StarsRate.svelte";
import type { components } from "$lib/api/types";

interface Props {
  userId: components["schemas"]["uuid"];
  spotId: components["schemas"]["uuid"];
  rating: number;
  comment: string;
  createdAt: string;
}
let spot: Props = $props();
</script>

<div class="bg-white flex justify-between items-center flex-col rounded-xl py-6 px-4 m-4 font-bold text-xl shadow-2xl space-y-2 w-full dark:bg-gray-900/95">
  <h1 class="pr-5 text-3xl">{spot.userId}</h1>
  <small class="text-xs font-thin">{spot.createdAt}</small>
  <div class="">
    <StarsRate star={spot.rating} />
  </div>
  <p class="text-sm">{spot.comment}</p>
</div>
