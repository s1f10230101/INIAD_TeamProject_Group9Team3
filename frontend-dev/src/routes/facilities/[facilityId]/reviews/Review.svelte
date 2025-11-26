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

<h1 class="pr-5 text-3xl">{spot.userId}</h1>
<small class="text-xs text-gray-700 font-thin">{spot.createdAt}</small>
<div class="">
  <StarsRate star={spot.rating} />
</div>
<p class="text-sm">{spot.comment}</p>
