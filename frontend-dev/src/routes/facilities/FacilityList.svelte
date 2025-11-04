<script lang="ts">
import type { components } from "$lib/types/api";

const setStarWidth = (node: HTMLElement, rating: number) => {
  const roundReview = Math.round(rating * 10) / 10;
  const widthPercentage = roundReview * 20;
  node.style.setProperty("--starWidth", `${widthPercentage}%`);
};

interface Prop {
  spot: {
    id: components["schemas"]["uuid"];
    name: string;
    description: string;
    address: string;
    createdAt: string;
  };
  averageRating: number;
  commentCount: number;
}
let { facilities }: { facilities: Prop[] } = $props();
</script>

{#each facilities as facility (facility.spot.id)}
  <div class="box">
    <a
      href={`facilities/reviews/${facility.spot.id}?rating=${facility.averageRating ?? 0}&count=${facility.commentCount ?? 0}`}
      class="facility-item"
      aria-label="{facility.spot.name}のレビューを見る"
    >
      <div class="line">
        <div class="item-info">
          <span class="facility-name" style="padding-right:1vw ;"
            >{facility.spot.name}</span
          >
          <span class="facility-location">{facility.spot.address}</span>
        </div>

        <div class="card-review_star">
          <span
            class="stars-clip"
            use:setStarWidth={facility.averageRating ?? 0}
          ></span>
          <span class="rating-value">{facility.averageRating ?? 0}</span>
          <span class="comment-count" style="padding-left:1vw ;"
            >コメント数({facility.commentCount})</span
          >
        </div>
      </div>
    </a>
  </div>
{/each}

<style>
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
    color: #5c4033;
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
    content: "★★★★★";
    position: absolute;
    width: 100%;
    top: 0;
    left: 0;
    letter-spacing: 0; /* 星がくっつく場合があるので調整 */
    color: #cccccc; /* 灰色（ベース） */
  }

  .stars-clip::after {
    content: "★★★★★";
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
