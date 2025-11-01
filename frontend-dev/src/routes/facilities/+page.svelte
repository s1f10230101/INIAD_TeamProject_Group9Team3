<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";
//import { facilities, type Facility } from "$lib/data/facilities";

//const data: Facility[] = facilities;

import { onMount } from "svelte";
import type { components } from "$lib/types/api";
import client from "$lib/api/client";

// ratingが含まれない場合の施設の変数の型
// let facilities: components["schemas"]["SpotResponse"][] = [];

// rating, commentCountを含まれるようにSpotWithRating型を定義した。
type SpotWithRating = components["schemas"]["SpotResponse"] & {
    averageRating?: number;
    commentCount?: number;
};
let facilities: SpotWithRating[] = [];
let isLoading = true;
let error: Error | null = null;

onMount(async () => {
    try {
        const { data: spotsData, error: spotsError } =
            await client.GET("/spots");
        if (spotsError) {
            throw new Error(
                spotsError.message || "施設一覧を取得することが失敗した",
            );
        }
        if (!spotsData) {
            facilities = [];
            return;
        }

        const facilitiesWithRatings: SpotWithRating[] = await Promise.all(
            spotsData.map(async (spot) => {
                const { data: reviewsData, error: reviewsError } =
                    await client.GET("/spots/{spotId}/reviews", {
                        params: { path: { spotId: spot.id } },
                    });
                let averageRating: number | undefined;
                let commentCount: number | undefined;
                if (reviewsError) {
                    console.warn(
                        `施設 ${spot.name} (${spot.id})のレビューを取得できませんでした`,
                        reviewsError,
                    );
                } else if (reviewsData && reviewsData.length > 0) {
                    const totalRating = reviewsData.reduce(
                        (sum, review) => sum + review.rating,
                        0,
                    );
                    averageRating = parseFloat(
                        (totalRating / reviewsData.length).toFixed(1),
                    );
                    commentCount = reviewsData.length;
                }

                return {
                    ...spot,
                    averageRating,
                    commentCount,
                };
            }),
        );
        facilities = facilitiesWithRatings;
    } catch (e) {
        if (e instanceof Error) {
            error = e;
        } else if (
            typeof e === "object" &&
            e !== null &&
            "message" in e &&
            typeof e.message === "string"
        ) {
            error = new Error(e.message);
        } else {
            error = new Error("予期しないエラーが起きました。");
        }
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
        <div class="review-main-container">
            <!--{#each data as facility (facility.id)}-->
            {#each facilities as facility (facility.id)}
            <div class="box">
                <a 
                    href={`facilities/reviews/${facility.id}`} 
                    class="facility-item"
                    aria-label="{facility.name}のレビューを見る"
                >   
                <div class="line">
                    <div class="item-info">
                        <span class="facility-name" style="padding-right:1vw ;">{facility.name}</span>
                        <!--<span class="facility-location">{facility.location}</span>-->
                        <span class="facility-location">{facility.address}</span>
                    </div>
                    
                    <div class="card-review_star" >
                        <span class="stars-clip" use:setStarWidth={facility.averageRating ?? 0}></span>
                        <!--<span class="rating-value">{facility.rating.toFixed(1)}</span>-->
                        <span class="rating-value">{facility.averageRating ?? 0}</span>
                        <span class="comment-count" style="padding-left:1vw ;">コメント数({facility.commentCount})</span>
                    </div>
                </div>
                </a>
            </div>
        {/each}
        </div>
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