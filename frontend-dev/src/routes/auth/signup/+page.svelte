<script lang="ts">
import { onMount } from "svelte";

import back10 from "$lib/assets/back10.png";
import back11 from "$lib/assets/back11.png";
import back12 from "$lib/assets/back12.png";
import back13 from "$lib/assets/back13.png";

import "../app.css";
import favicon from "$lib/assets/icon4.png";
import Header from "$lib/components/Header.svelte";
import Footer from "$lib/components/Footer.svelte";

let { children } = $props();

const backgroundImages = [back10, back11, back12, back13];

// 3. 状態変数を定義
let currentImage = $state(backgroundImages[0]);
let nextImage = $state(backgroundImages[1 % backgroundImages.length]);
let currentIndex = 0; // これはHTMLで使わないので `let` のままでOK
let isFadingOut = $state(false);

// 4. onMountでタイマーをセット
onMount(() => {
    const interval = setInterval(() => {
        isFadingOut = true;
        setTimeout(() => {
            currentIndex = (currentIndex + 1) % backgroundImages.length;
            currentImage = backgroundImages[currentIndex];
            nextImage =
                backgroundImages[(currentIndex + 1) % backgroundImages.length];
            isFadingOut = false;
        }, 1500); // CSSのアニメーション時間
    }, 7000); // 切り替え間隔

    return () => {
        clearInterval(interval);
    };
});

// SvelteKitのナビゲーション機能をインポート
import { goto } from "$app/navigation";

// ユーザー入力をバインドするための変数
let mailAddress = "";
let password = "";

// 登録処理（ダミー）
function handleLogin() {
    console.log("登録試行:", { mailAddress, password });

    // 実際のログイン処理はバックエンドで

    alert("登録に成功しました！（ダミー処理）");
    goto("/auth/login"); // ログインページに遷移
}
</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main class="center-content">
        <h1>新規登録画面</h1>

        <div class="register-container">
            <form>
                <div class="form-group">
                    <label for="mailAddress">メールアドレス</label>
                    <input type="email" id="mail-address" name="mailAddress" placeholder="メールアドレスを入力" bind:value={mailAddress} required>
                </div>

                <div class="form-group">
                    <label for="password">パスワード</label>
                    <input type="password" id="password" name="password" placeholder="パスワードを入力" bind:value={password} required>
                </div>

                <div class="button-group">
                    <button on:click={handleLogin} type="submit" class="btn-submit">ログイン</button>
                </div>
            </form>
        </div>

    </main>
</div>

<style>
    .full-screen-background{
        background-size: cover;
        background-position: center center; 
        background-attachment: fixed;
        background-image: var(--background-url); 
        display: flex;
        justify-content: center; /* 横方向の中央寄せ */
        align-items: center;   /* 縦方向の中央寄せ */
        padding-bottom: 10vw;
    }

    .center-content {
        padding: 5vw;
        padding-top: 15vw;
    }

    .center-content h1 {
        font-weight: bold;
        color: #5C4033;
        font-size: 28px;
        display: flex;
        justify-content: center; /* 横方向の中央寄せ */
        padding-bottom: 1vw;
        text-shadow: 1px 1px 5px #ffffff;
    }

/*  フォームコンテナの設定 */
    .register-container {
        width: 60vw; /* フォーム全体の幅 */
        padding: 3vh 3vw;
        background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
        border-radius: 10px; /* 角丸 */
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1); /* 軽い影 */
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
        color: #5C4033;
        font-size: 21px;
    }

    .form-group input[type="email"], 
    .form-group input[type="password"] {
        width: 100%;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 5px;
        box-sizing: border-box; /* paddingを含めた幅にする */
        font-size: 16px;
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
        background-color: #6D4C41;
        color: white;
        box-shadow: 0 3px 0 #4E342E;
    }

    .btn-submit:hover {
        background-color: #5D4037;
    }


</style>