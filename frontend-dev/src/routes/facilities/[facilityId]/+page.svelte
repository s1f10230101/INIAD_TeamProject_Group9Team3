<script lang="ts">
  import type { PageProps } from "./$types";
    import StarsRate from "$lib/components/StarsRate.svelte";

  let { params, data }: PageProps = $props();
  const facilityId: string = params.facilityId;
  const spotPromise = data.spotPromise;
  const reviewsPromise = data.reviewPromise;

  // レビュー平均値を通信が終わり次第thenメソッドチェーンで代入する
  let [averageRating, commentCount] = $state([0, 0]);
  if (reviewsPromise) {
    reviewsPromise
      .then((res) => res.data)
      .then((data) => {
        if (!data || data.length === 0) return;
        else {
          averageRating = parseFloat(
            (
              data.reduce((acc, review) => acc + review.rating, 0) / data.length
            ).toFixed(1),
          );
          commentCount = data.length;
        }
      });
  }
</script>

<div class="p-2 flex justify-center items-center flex-col">
  <h1 class="text-[#5c4033] text-4xl font-bold py-2">詳細ページ</h1>
  <div
    class="bg-white flex justify-between items-center flex-col rounded-xl py-6 px-4 m-4 text-[#5c4033] font-bold text-xl shadow-2xl space-y-2"
  >
    {#await spotPromise}
      <p>読み込み中...</p>
    {:then spotValue}
      <h1 class="pr-5 text-3xl">{spotValue.data?.name}</h1>
      <small class="text-xs text-gray-700 font-thin"
        >{spotValue.data?.address}</small
      >
      <StarsRate star={averageRating} />
      <p class="text-sm">{spotValue.data?.description}</p>
      <a
        href="/facilities/{facilityId}/reviews"
        class="text-sm font-normal">コメント数({commentCount})</a
      >
    {:catch spotError}
      <h1>エラー</h1>
      <p>施設ID: {facilityId} に対応する施設が見つかりませんでした。</p>
      {spotError}
    {/await}
  </div>

  <hr class="mb-4 border-[#5c4033]/50 w-full" />
  <a href="/facilities">施設一覧に戻る</a>
</div>
