<script lang="ts">
/**
 * @prop {star}
 */
// star is int 1~5
let { star }: { star: number } = $props();
</script>

<div class="flex justify-center gap-1 w-auto shrink-0 text-4xl">
  <!-- w-とh-で手動で幅を調整して奇跡的なかみ合いをしている。文字サイズを変更するとすぐ壊れる -->
  <span
    class="relative inline-block w-44 star-clips
    before:content-['★★★★★'] before: text-gray-300
    after:content-['★★★★★'] after:absolute after:left-0
    after:overflow-hidden after:text-yellow-300"
    style="--starWidth: {star * 20 + 1.5}%"
  ></span>
  <!-- なぜか1.5プラスすると数値に正確な幅になる -->
  <span class="text-orange-700 font-bold">{star.toFixed(2)}</span>
</div>

<style>
  .star-clips::after {
    width: var(--starWidth);
  }
</style>
