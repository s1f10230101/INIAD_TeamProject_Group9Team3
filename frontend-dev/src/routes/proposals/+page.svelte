<script lang="ts">
import client, { streamingRecvHelper } from "$lib/api/client";
import { type PageProps, type SubmitFunction } from "./$types";
import { enhance } from "$app/forms";

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

<div class="p-10 flex flex-col space-y-7">
  <h1
    class="font-bold text-[#5c4033] text-[1.75rem] flex"
  >
    体験したい旅行体験をご自由にお書きください
  </h1>
  <form method="POST" use:enhance={enhanceOption}>
    <div class="pb-4">
      <input
        type="text"
        bind:value={prompt}
        class="bg-white/90 p-3 rounded-xl w-full"
        name="prompt"
        placeholder="例：家族で温泉旅行"
      />
    </div>
    <p>{errorMsg}</p>
    <div>
      <button
        type="submit"
        class="w-full text-white bg-[#5c4033] rounded-3xl p-3 disabled:bg-[#9c877f] hover:bg-[#5c4033] hover:shadow-lg hover:translate-y-[-2px]"
        disabled={isLoading || !prompt}
      >
        {isLoading ? "生成中..." : "送信"}
      </button>
    </div>
  </form>
  <div class="mt-1 w-full">
    <div
      class="p-5 min-h-96 bg-[#FFF7E6] text-base text-[#5c4033] whitespace-pre-wrap rounded-4xl"
    >
      {#if form}
        <!-- JS無効の環境ではこちらが実行される -->
        {form.res}
      {:else}
        <!-- JS有効ならこちらがストリーム更新される -->
        {aiResponse}
      {/if}
    </div>
  </div>
</div>
