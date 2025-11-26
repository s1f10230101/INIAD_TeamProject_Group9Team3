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
  class="p-2 m-4 bg-gray-50/85 rounded-2xl shadow-sm text-[#5c4033] font-bold text-xl hover:shadow-2xl"
>
  <a
    href={`facilities/${facility.spot.id}`}
    aria-label="{facility.spot.name}のレビューを見る"
  >
    <div class="flex justify-between flex-col items-center space-y-3">
      <h3 class="text-3xl">{facility.spot.name}</h3>
      <span class="text-xs font-normal">{facility.spot.address}</span>

      <StarsRate star={facility.averageRating ?? 0} />
      <span class="flex justify-center text-base"
        >コメント数({facility.commentCount})</span
      >
    </div>
  </a>
</div>
