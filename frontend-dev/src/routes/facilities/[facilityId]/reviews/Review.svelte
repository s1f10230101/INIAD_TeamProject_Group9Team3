<script lang="ts">
  import type { components } from "$lib/types/api";

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
  <span class=""></span>
  <span class="">{spot.rating}</span>
</div>
<p class="text-sm">{spot.comment}</p>
