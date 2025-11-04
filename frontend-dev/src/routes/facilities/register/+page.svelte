<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";
import type { PageProps } from "./$types";
let { form }: PageProps = $props();
</script>

<div
  class="full-screen-background"
  style="--background-url: url('{backgroundImage}')"
>
  <main class="center-content">
    <div class="register-container">
      <h1>施設情報登録</h1>
      <form method="POST">
        <div class="form-group">
          <label for="facility-name">施設名</label>
          <input
            type="text"
            id="facility-name"
            name="facilityName"
            placeholder="施設名を入力"
            value={form?.facilityName}
          />
          {#if form?.missingName}
            <p class="text-red-600">名前は必須です</p>
          {/if}
        </div>

        <div class="form-group">
          <label for="location">所在地</label>
          <input
            type="text"
            id="location"
            name="location"
            placeholder="住所を入力"
            value={form?.location}
          />
        </div>

        <div class="form-group">
          <label for="hours">営業時間</label>
          <input
            type="text"
            id="hours"
            name="hours"
            placeholder="例: 9:00〜18:00"
            required
          />
        </div>

        <div class="form-group">
          <label for="price-range">価格帯</label>
          <input
            type="text"
            id="price-range"
            name="priceRange"
            placeholder="例: ¥1,000〜¥3,000"
            required
          />
        </div>

        <div class="form-group">
          <label for="description">概要・説明</label>
          <textarea
            id="description"
            name="description"
            placeholder="施設の特徴や魅力を詳しく記入してください"
            rows="5"
          >{form?.description}</textarea>
        </div>

        <div class="button-group">
          <button type="submit" class="btn-submit">施設情報を登録する</button>
        </div>
      </form>
    </div>
  </main>
</div>

<style>
  .full-screen-background {
    background-size: cover;
    background-position: center center;
    background-attachment: fixed;
    background-image: var(--background-url);
    display: flex;
    justify-content: center; /* 横方向の中央寄せ */
    align-items: center; /* 縦方向の中央寄せ */
  }

  .center-content {
    padding: 5vw;
  }

  /*  フォームコンテナの設定 */
  .register-container {
    width: 60vw; /* フォーム全体の幅 */
    padding: 3vh 3vw;
    background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
    border-radius: 10px; /* 角丸 */
    box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1); /* 軽い影 */
  }

  .register-container h1 {
    font-weight: bold;
    color: #5c4033;
    font-size: 28px;
  }

  /* -------------------------------------- */
  /*  フォーム部品のスタイル */
  /* -------------------------------------- */
  .form-group {
    margin-bottom: 25px;
  }

  .form-group label {
    display: block;
    margin-bottom: 8px;
    font-weight: bold;
    color: #5c4033;
    font-size: 21px;
  }

  .form-group input[type="text"],
  .form-group textarea {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    box-sizing: border-box;
    font-size: 16px;
    background-color: #fcfcfc;
    resize: vertical; /* 💡 textareaのリサイズを縦方向のみに制限 */
  }

  /* -------------------------------------- */
  /*  ボタンのスタイル */
  /* -------------------------------------- */
  .button-group {
    margin-top: 40px;
  }

  .button-group button {
    display: block;
    width: 100%;
    padding: 12px;
    margin-bottom: 15px;
    border: none;
    border-radius: 5px;
    font-size: 18px;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s;
  }

  /* 登録ボタン (濃い茶色) */
  .btn-submit {
    background-color: #6d4c41;
    color: white;
    box-shadow: 0 3px 0 #4e342e;
  }

  .btn-submit:hover {
    background-color: #5d4037;
  }
</style>
