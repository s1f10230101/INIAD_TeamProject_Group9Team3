<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";

// SvelteKitから現在のページの情報を取得
import { page } from "$app/stores";

// 元の施設データと型をインポート
import { facilities, type Facility } from "$lib/data/facilities";

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
</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main class="center-content">
        {#if facilityData}
        <h1>レビューを投稿する</h1>
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
    
    <a href="/facilities">一覧に戻る</a>

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



</style>