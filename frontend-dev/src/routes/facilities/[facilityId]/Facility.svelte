<script lang="ts">
import type { components } from "$lib/api/types";
import StarsRate from "$lib/components/StarsRate.svelte";

interface Prop {
  spot: {
    id: components["schemas"]["uuid"];
    name: string;
    description: string;
    address: string;
    createdAt: string;
  };
  averageRating: number;
  commentCount: number;
}
let facility: Prop = $props();
</script>

<div
  class="p-3 rounded-3xl text-[#5c4033] font-bold text-xl w-full bg-gray-50/85"
>
  <a
    href="/facilities/{facility.spot.id}/reviews"
    aria-label="{facility.spot.name}のレビューを見る"
  >
    <div class="flex justify-between flex-col items-center space-y-3">
      <h3 class="text-3xl">{facility.spot.name}</h3>
      <span class="text-xs font-normal">{facility.spot.address}</span>

      <StarsRate star={facility.averageRating ?? 0} />
      <p class="text-sm">{facility.spot.description}</p>
      <span class="flex justify-center text-base"
        >コメント数({facility.commentCount})</span
      >
    </div>
  </a>
</div>
