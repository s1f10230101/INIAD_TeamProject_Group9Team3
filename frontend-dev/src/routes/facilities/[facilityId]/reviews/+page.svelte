<script lang="ts">
  import type { PageProps } from "./$types";
  import Review from "./Review.svelte";
  let { data, params }: PageProps = $props();
  const { reviewPromise } = data;
</script>

<div class="container mx-auto p2 flex justify-center items-center flex-col">
  <h1 class="text-[#5c4033] text-4xl font-bold py-2">レビュー</h1>
  <hr class="mb-4 border-[#5c4033]/50 w-full" />
  <a href="./reviews/register" class="bg-blue-600 w-full text-center p-3 rounded-2xl text-white"
    >+コメントを書き込む</a
  >
  <div
    class="bg-white flex justify-between items-center flex-col rounded-xl py-6 px-4 m-4 text-[#5c4033] font-bold text-xl shadow-2xl space-y-2"
  >
    {#await reviewPromise}
      <p>読み込み中...</p>
    {:then review}
      {#if review.data}
        {#each review.data.filter((e) => e !== undefined) as spot}
          <Review {...spot} />
        {/each}
      {:else}
        <p class="text-red-700">
          {review.error.message}
        </p>
      {/if}
    {:catch spotError}
      <h1>エラー</h1>
      <p>対応する施設が見つかりませんでした。</p>
      {spotError}
    {/await}
  </div>
</div>
