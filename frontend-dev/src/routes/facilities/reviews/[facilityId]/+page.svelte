<script lang="ts">
	import backgroundImage from "$lib/assets/back10.png";
	import { page } from "$app/stores";
	import { onMount } from "svelte";
	import client from "$lib/api/client";
	import type { components } from "$lib/types/api";

	type Spot = components["schemas"]["SpotResponse"];
	type Review = components["schemas"]["ReviewResponse"];

	const facilityId = $page.params.facilityId;

	let facilityData: Spot | null = null;
	let reviews: Review[] = [];
	let isLoading = true;
	let error: string | null = null;

	onMount(async () => {
		isLoading = true;
		try {
			// 1. 施設の詳細情報を取得
			const { data: spotData, error: spotError } = await client.GET("/spots/{spotId}", {
				params: { path: { spotId: facilityId } },
			});

			if (spotError) throw new Error("施設の情報の取得に失敗しました。");
			facilityData = spotData;

			// 2. レビューの一覧を取得
			const { data: reviewsData, error: reviewsError } = await client.GET("/spots/{spotId}/reviews", {
				params: { path: { spotId: facilityId } },
			});

			if (reviewsError) throw new Error("レビューの取得に失敗しました。");
			reviews = reviewsData || [];

		} catch (e: any) {
			error = e.message;
		} finally {
			isLoading = false;
		}
	});

	const setStarWidth = (node: HTMLElement, rating: number) => {
		const roundReview = Math.round(rating * 10) / 10;
		const widthPercentage = roundReview * 20;
		node.style.setProperty("--starWidth", `${widthPercentage}%`);
	};

</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main class="center-content">
        {#if isLoading}
            <p>読み込み中...</p>
        {:else if error}
            <h1>エラー</h1>
            <p>{error}</p>
        {:else if facilityData}
            <h1>{facilityData.name} のレビュー</h1>
            <hr class="custom-line">

            <!-- レビュー一覧 -->
            <div class="reviews-list">
                {#each reviews as review}
                    <div class="review-card">
                        <div class="card-review_star" >
                            <span class="stars-clip" use:setStarWidth={review.rating}></span>                    
                            <span class="rating-value">{review.rating.toFixed(1)}</span>
                        </div>
                        <p class="review-comment">{review.comment}</p>
                        <p class="review-meta">投稿者: {review.userId.substring(0, 8)}... | 投稿日: {new Date(review.createdAt).toLocaleDateString()}</p>
                    </div>
                {:else}
                    <p>この施設にはまだレビューがありません。</p>
                {/each}
            </div>
        {/if}
        
        <a href="/facilities">施設一覧に戻る</a>
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