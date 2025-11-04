<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";
import client, { streamingRecvHelper } from "$lib/api/client";
import { type PageProps } from "./$types";
import { enhance } from "$app/forms";

import { type SubmitFunction } from "@sveltejs/kit";

let { form }: PageProps = $props();

let aiResponse = $state("");
let prompt = $state("");
let isLoading = $state(false);

// ブラウザでjavascriptが有効なときはストリーム受信するためのenhanceオプション
const option: SubmitFunction = async ({ formData, cancel }) => {
  cancel();
  const promptData = formData.get("prompt");
  if (!promptData) return;
  const prompt = promptData.toString();

  isLoading = true;
  const { response } = await client.POST("/plans", {
    body: {
      prompt: prompt,
    },
    parseAs: "stream",
  });
  await streamingRecvHelper(response, (recvText) => {
    aiResponse += recvText;
  });

  return async ({ update }) => {
    await update({ reset: true });
    isLoading = false;
  };
};
</script>

<div
  class="full-screen-background"
  style="--background-url: url('{backgroundImage}')"
>
  <main class="center-content">
    <h1>体験したい旅行体験をご自由にお書きください</h1>
    <form method="POST" use:enhance={option}>
      <input
        type="text"
        bind:value={prompt}
        class="form"
        name="prompt"
        placeholder="例：家族で温泉旅行"
      />
      <button type="submit" class="submit-button" disabled={isLoading}>
        {isLoading ? "生成中..." : "送信"}
      </button>
    </form>
    <div class="ai-response-container">
      <div class="ai-response-box">
        {#if form}
          <!-- JS無効の環境ではこちらが実行される -->
          {form.res}
        {:else}
          <!-- JS有効ならこちらがストリーム更新される -->
          {aiResponse}
        {/if}
      </div>
    </div>
  </main>
</div>

<style>
  .full-screen-background {
    background-size: cover;
    background-position: center center;
    background-attachment: fixed;
    background-image: var(--background-url);
    padding-top: 10vw;
    display: flex;
    justify-content: center; /* 横方向の中央寄せ */
    align-items: center; /* 縦方向の中央寄せ */
  }

  .center-content {
    padding: 5vw;
  }

  .center-content h1 {
    font-weight: bold;
    color: #5c4033;
    font-size: 28px;
    display: flex;
    justify-content: center; /* 横方向の中央寄せ */
    padding-bottom: 1vw;
    text-shadow: 1px 1px 5px #ffffff;
  }

  .form {
    width: 60vw; /* フォーム全体の幅 */
    padding: 1.5vh 1vw;
    background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
    border-radius: 10px; /* 角丸 */
    box-shadow: 0 4px 15px #5c4033; /* 軽い影 */
    margin-bottom: 30px;
  }

  .ai-response-container {
    width: 60vw;
    margin-top: 20px;
  }

  .ai-response-box {
    width: 100%;
    min-height: 20vh; /* 最低の高さを確保 */
    padding: 20px;
    background-color: rgba(255, 255, 255, 0.95); /* より不透明な白背景 */
    border-radius: 10px;
    box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
    font-size: 16px;
    line-height: 1.6;
    color: #5c4033;
    white-space: pre-wrap; /* テキストの改行を有効にする */
    padding-bottom: 20vw;
  }
</style>
