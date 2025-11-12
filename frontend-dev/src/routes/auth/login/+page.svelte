<script>

// SvelteKitのナビゲーション機能をインポート
import { goto } from "$app/navigation";

// ユーザー入力をバインドするための変数
let mailAddress = "";
let password = "";

// ログイン処理（ダミー）
function handleLogin() {
    console.log("ログイン試行:", { mailAddress, password });

    // 実際のログイン処理はバックエンドで

    alert("ログインに成功しました！（ダミー処理）");
    goto("/"); // トップページ（ルート）に遷移
}

// 新規登録画面への遷移
function handleSignup() {
    // /register 画面に遷移
    goto("/auth/signup");
}
</script>
<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main class="center-content">
        <h1>ログイン画面</h1>

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

        <div class="button-group">
            <button on:click={handleSignup} type="submit" class="btn-submit">新規登録</button>
        </div>

        <div class="button-group">
            <button on:click={handleSignup} type="submit" class="btn-submit">新規登録</button>
        </div>
    </main>
</div> 


<style>
    .full-screen-background{
        display: flex;
        justify-content: center; /* 横方向の中央寄せ */
        align-items: center;   /* 縦方向の中央寄せ */
        padding-bottom: 5vw;
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