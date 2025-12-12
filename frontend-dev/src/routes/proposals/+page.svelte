<script lang="ts">
import client, { streamingRecvHelper } from "$lib/api/client";
import { escapeHtml } from "$lib/index";
import { type PageProps, type SubmitFunction } from "./$types";
import { enhance } from "$app/forms";
import { marked } from "marked";

let { form }: PageProps = $props();

let errorMsg = $state(form ? form.invaild : "");
let aiResponse = $state("");
let prompt = $state("");
let isLoading = $state(false);

// ブラウザでjavascriptが有効なときはストリーム受信するためのenhanceオプション
const enhanceOption: SubmitFunction = async ({
  formData,
  cancel,
  formElement,
}) => {
  cancel(); // フォーム送信をキャンセル(サーバーでストリームは使えないのでJSオンの時はサーバー処理はしない)
  const promptData = formData.get("prompt");
  if (!promptData) {
    errorMsg = "err";
    return;
  }
  const prompt = promptData.toString();
  formElement.reset();

  aiResponse = "";
  isLoading = true;
  const { response } = await client.POST("/plans", {
    body: { prompt: prompt },
    parseAs: "stream",
  });
  await streamingRecvHelper(response, (recvText) => {
    aiResponse += recvText;
  });

  isLoading = false;
};
</script>

<div class="pt-5 space-y-7 w-full">
  <h1 class="font-bold text-3xl text-center">
    体験したい旅行体験をご自由にお書きください
  </h1>
  <div class="p-5 min-h-96 bg-amber-50/90 rounded-4xl dark:bg-gray-900/90">
    <!-- prose lg:prose-xlは、tailwindのリセットcssを部分的に無効化;MarkDownを正しく表示するため -->
    <!-- TODO: サニタイズ -->
    <article class="prose wrap-break-word lg:prose dark:prose-invert">
      {#if form}
        <!-- JS無効の環境ではこちらが実行される -->
        {@html marked.parse(form.res!)}
      {:else}
        <!-- JS有効ならこちらがストリーム更新される -->
        {@html marked.parse(escapeHtml(aiResponse))}
      {/if}
    </article>
  </div>
  <form
    method="POST"
    class="fixed bottom-0 right-0 left-0 pb-10 z-2 flex flex-wrap container p-2 m-auto w-full justify-end"
    use:enhance={enhanceOption}
  >
    <div
      class="bg-white/95 flex-12 rounded-xl py-1
      dark:bg-gray-700/95 flex flex-col justify-center"
    >
      <input
        type="text"
        bind:value={prompt}
        class="p-2 w-full m-auto wrap-break-word border-none
          row-span-1 col-span-12 focus:outline-0"
        name="prompt"
        placeholder="例：家族で温泉旅行"
      />
      <div class="justify-end w-full flex p-1 pr-2">
        <button
          type="submit"
          class="text-white text-xl mt-2 disabled:bg-primary-light-300
          bg-primary-light-500 hover:shadow-lg hover:-translate-y-0.5
          rounded-full w-10 h-10 justify-end"
          disabled={isLoading || !prompt}
          aria-label="AIに送信"
        >
          ↑
        </button>
      </div>
      <p class="text-red-700">{errorMsg}</p>
    </div>
  </form>
</div>
