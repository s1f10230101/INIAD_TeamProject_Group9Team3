<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";

// SvelteKitから現在のページの情報を取得
import { page } from "$app/stores";

// 元の施設データと型をインポート
import { facilities, type Facility } from "$lib/data/facilities";

import { goto } from '$app/navigation';

// 1. URLのパラメータ（[facilityId]の部分）を取得し、数値に変換
// $page.params.facilityId は文字列なので、Number()で数値化します。
const facilityId = Number($page.params.facilityId);

// 2. 施設IDに基づいて、該当する施設データを検索
const facilityData: Facility | undefined = facilities.find(
    (f) => f.id === facilityId,
);

const setStarWidth = (node: HTMLElement, rating: number) => {
    const roundReview = Math.round(rating * 10) / 10;
    const widthPercentage = roundReview * 20;
    node.style.setProperty("--starWidth", `${widthPercentage}%`);
};

let isDetailVisible: boolean = false;

// 詳細表示の状態を切り替える関数
const toggleDetail = () => {
    isDetailVisible = !isDetailVisible;
};

const updateStarWidth = (node: HTMLElement, rating: number) => {
    // 値が更新されるたびに実行される関数を返す
    return {
        update(newRating: number) {
            const roundReview = Math.round(newRating * 10) / 10;
            const widthPercentage = roundReview * 20;
            node.style.setProperty("--starWidth", `${widthPercentage}%`);
        }
    };
};

// フォームデータ保持用の変数
let reviewTitle: string = "";
let reviewContent: string = "";
let ratingValue: number = 0.0; // 既存の変数を使用

// 画面状態を管理する変数 (false: 入力画面, true: 確認画面)
let isConfirmMode: boolean = false;

// 1. 確認ボタンが押された時の処理
const handleConfirm = (event: Event) => {
    // フォームのデフォルト送信（ページ遷移）を防止
    event.preventDefault(); 
    
    if (!reviewTitle || !reviewContent || ratingValue === 0.0) {
        return; 
    }
    
    isConfirmMode = true;
};

// 2. 投稿ボタンが押された時の処理
const handleSubmit = () => {
    // ★ ここに投稿データをサーバーに送信するロジックを実装 ★
    const reviewData = {
        facilityId: facilityId,
        title: reviewTitle,
        content: reviewContent,
        rating: ratingValue,
    };
    
    console.log("投稿データ:", reviewData);
    
    // 投稿後の処理（例: 一覧ページへ移動、完了メッセージ表示など）
    alert("レビューを投稿しました！");

    goto('/facilities');
};

// 3. 修正ボタンが押された時の処理
const handleEdit = () => {
    // 入力画面に戻る
    isConfirmMode = false;
};
</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main class="center-content">
        {#if facilityData}
        <h1>レビューを投稿する</h1>
        <div class="progress-bar-container">
            <div 
                class="progress-step" 
                class:is-active={!isConfirmMode} 
                class:is-done={isConfirmMode}
            >
                1.情報入力
            </div>
            <div 
                class="progress-step" 
                class:is-active={isConfirmMode}
            >
                2.確認
            </div>
            <div class="progress-step"> 
                3.完了
            </div>
        </div>
        <hr class="custom-line">
        <div class="review-main-container">
            <button on:click={toggleDetail}>
                <div class="line box">
                    <div class="item-info">
                        <span class="facility-name" style="padding-right:1vw ;">{facilityData.name}</span>
                        <span class="facility-location">{facilityData.location}</span>
                    </div>
                    
                    <div class="card-review_star" >
                        <span class="stars-clip" use:setStarWidth={facilityData.rating}></span>                    
                        <span class="rating-value">{facilityData.rating.toFixed(1)}</span>
                        <span class="comment-count" style="padding-left:1vw ;">コメント数({facilityData.commentCount})</span>
                    </div>
                </div>
            </button>
            {#if isDetailVisible} 
            <div class="container">
                <h1>{facilityData.name} のレビュー詳細</h1>
                <p>場所: {facilityData.location}</p>
                <p>営業時間: {facilityData.openHours}</p>
                <p>料金: {facilityData.price}</p>
                <p>説明: {facilityData.explanation}</p>
                
                <h2>評価</h2>
                <p>総合評価: {facilityData.rating.toFixed(1)} / コメント数: {facilityData.commentCount}件</p>
            </div>
            {/if}
        </div>
        {:else}
        <h1>エラー</h1>
        <p>施設ID: {facilityId} に対応する施設が見つかりませんでした。</p>
    {/if}
    
    <a href="/facilities" style="color: #5C4033; font-weight: bold; text-shadow: 5px 5px 5px #ffffff;">一覧に戻る</a>
    {#if !isConfirmMode}
    <div class="comment-container">
        <form class="review-form-grid" on:submit|preventDefault={handleConfirm}>
            <div class="form-row">
                <label for="review-title" class="grid-label">レビュータイトル</label>
                <input type="text" id="review-title" name="reviewTitle" class="grid-input" placeholder="タイトルを入力" required bind:value={reviewTitle}>
                <div class="grid-side-panel"></div>
            </div>

            <div class="form-row review-content-row">
                <label for="review-content" class="grid-label">レビュー内容</label>
                <textarea id="review-content" name="reviewContent" class="grid-input" placeholder="詳細を入力" required bind:value={reviewContent}></textarea>
                <div class="grid-side-panel"></div>
            </div>

            <div class="form-row rating-row">
                <label for="facility-rating" class="grid-label">施設評価</label>
                <div class="grid-input-content">
                <input type="number" id="facility-rating" name="facilityRating" class="rating-input"min="0.0" max="5.0" step="0.1" placeholder="0.0 〜 5.0" required bind:value={ratingValue}> 
                    <div class="card-review_star" >
                        <span class="stars-clip" use:updateStarWidth={ratingValue}></span>                    
                        <span class="rating-value"></span>
                    </div>
            </div>
                <div class="grid-side-panel"></div>
            </div>

            <div class="button-row">
                <button type="submit" class="btn-confirm">確認</button>
            </div>
        </form>
    </div>
    {:else}
    <div class="comment-container confirmation-view">
                <h2>以下の内容で投稿します。よろしいですか？</h2>
                
                <div class="confirmation-detail">
                    <p><strong>施設名:</strong> {facilityData?.name}</p>
                    <p><strong>タイトル:</strong> {reviewTitle}</p>
                    <p><strong>内容:</strong></p>
                    <pre>{reviewContent}</pre>
                    <p><strong>評価:</strong> 
                        <span class="card-review_star">
                            <span class="stars-clip" style={`--starWidth: ${(ratingValue / 5.0) * 100}%;`}></span>
                            <span class="rating-value">{ratingValue.toFixed(1)}</span>
                        </span>
                    </p>
                </div>

                <div class="button-row action-buttons">
                    <button class="btn-confirm" on:click={handleSubmit} style="margin-right: 15px; background-color: #388E3C;">投稿する</button>
                    <button class="btn-confirm" on:click={handleEdit} style="background-color: #795548;">修正する</button>
                </div>
            </div>
        {/if}

    </main>
</div>

<style>
    .full-screen-background{
        background-size: cover;
        background-position: center center; 
        background-attachment: fixed;
        background-image: var(--background-url); 
        padding-top: 5vw;
        display: flex;
        justify-content: center; /* 横方向の中央寄せ */
        align-items: center;   /* 縦方向の中央寄せ */
        padding-bottom: 5vw;
    }

    h1 {
        color: #5C4033;
        font-weight: bold;
        font-size: 28px; 
        text-shadow: 5px 5px 5px #ffffff;
    }

    h2 {
        color: #5C4033;
        font-weight: bold;
        font-size: 16px; 
    }

    .custom-line {
        border: 2px solid #5C4033;
        margin-bottom: 1rem;
    }

    .center-content {
        padding: 5vw;
    }

    .review-main-container {
        width: 60vw; /* フォーム全体の幅 */
        padding: 1.5vh 1vw;
        background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
        border-radius: 10px; /* 角丸 */
        box-shadow: 0 4px 15px #5C4033; /* 軽い影 */
        margin-bottom: 30px;
    }

    .comment-container {
        width: 60vw; /* フォーム全体の幅 */
        padding: 1.5vh 1vw;
        background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
        border-radius: 10px; /* 角丸 */
        box-shadow: 0 4px 15px #5C4033; /* 軽い影 */
        margin-bottom: 30px;
        margin-top: 30px;
    }

    .line {
        display: flex;
        justify-content: space-between;
    }


    .box {
        width: 56vw; /* フォーム全体の幅 */
        border-radius: 10px; /* 角丸 */
        box-shadow: 0 4px 15px #a3806f; /* 軽い影 */
        padding: 1.5vh 1vw;
        margin: 1vw;
        color: #5C4033;
        font-weight: bold;
        font-size: 20px; 

    }

    .box:hover {
        box-shadow: 0 4px 15px #e2ae3f;
    }

    .container {
        color: #5C4033;
        font-weight: bold;
        font-size: 20px; 
    }

    /* ★マークでレビュー表示 */

    .card-review_star {
        display: flex;
        align-items: center;
        gap: 5px;
        /* 星のサイズを固定し、Flexアイテムとしての挙動を安定させる */
        font-size: 20px; 
        width: auto; /* Flexアイテムとして幅は自動調整 */
        flex-shrink: 0; /* 縮小させない */
    }

    /* 塗りつぶしのための星の描画コンテナ */
    .stars-clip {
        /* starWidthが適用されるコンテナ */
        position: relative; /* 評価値と並べるための Flexbox と調和させるため */
        display: inline-block;
        
        /* 5つの星の幅を確保 (font-size: 20pxの場合) */
        width: 100px; /* 20px * 5 = 100px (例として固定) */
        height: 20px;
        line-height: 1em;
    }
    
    .stars-clip::before {
        content: '★★★★★';
        position: absolute;
        width: 100%;
        top: 0;
        left: 0;
        letter-spacing: 0; /* 星がくっつく場合があるので調整 */
        color: #CCCCCC; /* 灰色（ベース） */
    }

    .stars-clip::after {
        content: '★★★★★';
        position: absolute;
        top: 0;
        left: 0;
        
        /* 核心部分: Svelteから受け取った変数で幅を制御 */
        width: var(--starWidth); 
        
        /* グラデーションは不要。overflow: hidden で幅制御と色付けを両立させる */
        overflow: hidden;
        white-space: nowrap;
        color: #ffcf32; /* 金色（前景） */
    }

    .review-form-grid {
    /* Gridレイアウトを有効にする */
    display: grid;
    /* 💡 3つの列のサイズを定義 */
    grid-template-columns: auto 1fr 150px; /* ラベル幅 | 残り全て | 装飾パネル幅 */
    
    /* フォーム全体のパディングと角丸（外側のコンテナは別途必要） */
    padding: 20px;
    border: 1px solid #C4A59B;
    border-radius: 10px;
    background-color: #FFF;
}

.form-row {
    display: contents; 
}

/* ラベルセクション */
.grid-label {
    background-color: #FFF5EC; 
    padding: 10px 15px;
    border: 1px solid #C4A59B;
    color: #5C4033;
    /* Gridの子要素として中央揃え */
    align-self: stretch; 
    display: flex;
    align-items: center; 
    justify-content: flex-end; 
    border-radius: 10px; /* 角丸 */
    font-weight: bold;
}

/* 入力セクション */
.grid-input, .grid-input-content {
    background-color: white; 
    padding: 10px 15px;
    border-radius: 10px; /* 角丸 */
    border: 1px solid #C4A59B;
}

/* 装飾パネルセクション */
.grid-side-panel {
    background-color: #D3C4B8; /* 右の茶色いパネル背景 */
    border: 1px solid #C4A59B;
    border-radius: 10px; /* 角丸 */
}

.grid-input {
    width: 100%;
    border: 1px solid #C4A59B;
    box-sizing: border-box;
    border-radius: 10px; /* 角丸 */
}

/* レビュー内容 (Textarea) の高さ調整 */
.review-content-row .grid-label {
    align-items: flex-start; 
    padding-top: 20px;
}
.review-content-row textarea {
    min-height: 150px; 
    resize: none;
}

.rating-row .grid-input-content {
    /* 評価スターと数値の配置 */
    display: flex;
    align-items: center;
}

.button-row {
    grid-column: 1 / span 3;
    padding-top: 20px;
}
.btn-confirm {
    display: block;
    width: 100%;
    padding: 15px;
    background-color: #6D4C41; /* 濃い茶色 */
    color: white;
    border: none;
    border-radius: 5px;
    font-size: 18px;
    font-weight: bold;
}

/* 確認画面のコンテナ */
.confirmation-view {
    /* 確認画面の見た目を整える */
    padding: 40px;
    text-align: left;
    color: #5C4033;
}

.confirmation-detail {
    margin-top: 20px;
    padding: 15px;
    border: 1px solid #C4A59B;
    border-radius: 5px;
    background-color: #FFF5EC;
}

.confirmation-detail p {
    margin-bottom: 10px;
}

.confirmation-detail strong {
    display: inline-block;
    min-width: 80px;
    margin-right: 10px;
}

.confirmation-detail pre {
    white-space: pre-wrap; /* 内容の改行を保持 */
    font-family: inherit;
    margin-top: 5px;
    padding-left: 10px;
    border-left: 2px solid #D3C4B8;
}

.action-buttons {
    display: flex;
    justify-content: center;
    margin-top: 30px;
}

.action-buttons button {
    width: auto;
    min-width: 150px;
    flex-grow: 1;
}

.progress-bar-container {
    display: flex;
    justify-content: center; /* フォームに揃えて中央に配置 */
    width: 60vw; /* フォームと同じ幅に合わせる */
    margin-bottom: 20px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
}

.progress-step {
    flex-grow: 1;
    padding: 10px 30px 10px 20px;
    font-size: 16px;
    font-weight: bold;
    color: #5C4033; /* デフォルトの文字色 */
    background-color: #E0E0E0; /* デフォルトの背景色（灰色） */
    text-align: center;
    position: relative;
    /* 矢印の形状を作るために必須 */
    clip-path: polygon(0 0, calc(100% - 15px) 0, 100% 50%, calc(100% - 15px) 100%, 0 100%, 15px 50%);
    z-index: 10;
}

/* 1番目のステップは左端を四角にする */
.progress-bar-container .progress-step:first-child {
    padding-left: 15px;
    clip-path: polygon(0 0, calc(100% - 15px) 0, 100% 50%, calc(100% - 15px) 100%, 0 100%);
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
}

/* 最後のステップは右端を四角にする */
.progress-bar-container .progress-step:last-child {
    clip-path: polygon(0 0, 100% 0, 100% 100%, 0 100%, 15px 50%);
    padding-right: 15px;
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
}

/* -------------------------------------- */
/* ステータスの色分け */
/* -------------------------------------- */

/* デフォルトの非アクティブ状態 (灰色) */
.progress-step {
    background-color: #E0E0E0; 
    color: #5C4033;
    /* ... 他のスタイル ... */
}

/* 実行済み (完了) ステップのスタイル */
.progress-step.is-done {
    background-color: #A1887F; /* 例: 完了を示す薄い茶色 */
    color: white;
}

/* 現在アクティブなステップのスタイル */
.progress-step.is-active {
    background-color: #6D4C41; /* 現在地を示す濃い茶色 */
    color: white;
}


</style>