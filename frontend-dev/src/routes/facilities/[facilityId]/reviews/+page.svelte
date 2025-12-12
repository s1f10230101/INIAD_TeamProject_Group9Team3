<script lang="ts">
import type { PageProps } from "./$types";
import Review from "./Review.svelte";
let { data }: PageProps = $props();
const { reviewPromise } = data;
</script>

<div class="container mx-auto p2 flex justify-center items-center flex-col">
  <h1 class="text-4xl font-bold py-2">レビュー</h1>
  <hr class="mb-4 border-primary-light-500/50 w-full" />
  <a href="./reviews/register" class="bg-primary-light-500 w-full text-center p-3 rounded-2xl text-white hover:bg-primary-light-500 hover:shadow-lg hover:-translate-y-0.5"
    >+コメントを書き込む</a
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
