<script lang="ts">
  import ProgressBer from "./ProgressBer.svelte";
  import { type PageProps } from "./$types";
  import Textarea from "$lib/components/Textarea.svelte";
  import { enhance } from "$app/forms";
  import StarRateInput from "$lib/components/StarRateInput.svelte";
  import StarsRate from "$lib/components/StarsRate.svelte";
  let reviewContent = $state("");
  let { data, form }: PageProps = $props();
  let reviewRating = $state(1);
  let dialog: HTMLDialogElement;
</script>

<div class="flex items-center flex-col pt-6">
  <ProgressBer />
  <form class="" method="POST" use:enhance={()=>{
      dialog.close()
      return ({update})=>{update({invalidateAll: true, reset:true})}}
    }>
    <Textarea
      id="review-content"
      name="reviewContent"
      placeholder="詳細を入力"
      bind:value={reviewContent}
      label="レビュー内容"
      error={(form)? form.ContentError: ""}
    />
    <StarRateInput bind:star={reviewRating} name="star" />

    <button
      type="button"
      class="text-center w-full bg-blue-500/95 rounded-3xl p-4 hover:bg-blue-800/95"
      onclick={() => dialog.showModal()}>確認</button
    >

    <!-- 以下、確認ダイアログで開かれる要素 -->
    <dialog
      bind:this={dialog}
      closedby="any"
      class="fixed m-auto p-10 rounded-3xl backdrop:bg-gray-500/80"
    >
      <h2>以下の内容で投稿します。よろしいですか？</h2>

      <div class="confirmation-detail">
        {#if !data.spotData}
          <p class="text-red-600">{data.errorMsg}</p>
        {:else}
          <strong>施設名:</strong> <h3  class="text-3xl text-center">{data.spotData.name}</h3>
        {/if}

        <p><strong>内容:</strong></p>
        <pre>{reviewContent}</pre>
        <p>
          <strong>評価:</strong>
          <StarsRate star={reviewRating} />
        </p>
      </div>

      <div class="flex justify-center flex-col space-4 gap-4">
        <button type="submit" class=" bg-green-700 rounded-2xl p-2"
          >投稿する</button
        >
        <button
          type="button"
          class="bg-red-700 rounded-2xl p-2"
          onclick={() => dialog.close()}>修正する</button
        >
      </div>
    </dialog>
  </form>
</div>
