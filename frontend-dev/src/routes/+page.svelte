<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";
</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main>
        <div class="vertical-sections-wrapper1">

            <section class="vertical-text">
                <h1>あなただけの旅行体験</h1>
                <p>
                    本アプリは従来のガイドブックや<br />
                    SEOでは見つけられない<br />
                    まだ知られていない魅力的な場所と<br />
                    出会える旅行プラン検索アプリです
                </p>
            </section>

            <section class="vertical-text">
                <h1>隠れた名所を見つける</h1>
                <p>
                    地元の事業者さんが登録した<br />
                    こだわりの観光施設や体験を<br />
                    豊富にそろえています
                </p>
            </section>
        </div>

        <div class="vertical-sections-wrapper2">
            <section class="vertical-text">
                <h1>あなただけの旅行体験</h1>
                <p>
                    本アプリは従来のガイドブックや<br />
                    SEOでは見つけられない<br />
                    まだ知られていない魅力的な場所と<br />
                    出会える旅行プラン検索アプリです
                </p>
            </section>

            <section class="vertical-text">
                <h1>あなただけの旅行体験</h1>
                <p>
                    本アプリは従来のガイドブックや<br />
                    SEOでは見つけられない<br />
                    まだ知られていない魅力的な場所と<br />
                    出会える旅行プラン検索アプリです
                </p>
            </section>
        </div>

        <div class="horizontal-text">
            <h1>Trip-Canvasで、あなただけの新しい旅の物語を始めましょう</h1>
        </div>


    </main>
</div>

<style>
    .full-screen-background{
        background-size: cover;
        background-position: center center;
        background-attachment: fixed;
        background-image: var(--background-url);
    }

    .vertical-sections-wrapper1 {
        /* セクションを横並びにする */
        display: flex;
        /* セクション間の間隔を調整 */
        gap: 20.83vw;
        padding-bottom: 15vw;

        /* 縦書きのテキストが画面中央に配置されるように */
        align-items: center;
        justify-content: center;

    }

    .vertical-sections-wrapper2 {
        /* セクションを横並びにする */
        display: flex;
        /* セクション間の間隔を調整 */
        gap: 17.18vw;
        padding-bottom: 15vw;

        /* 縦書きのテキストが画面中央に配置されるように */
        align-items: center;
        justify-content: center;

    }

    .vertical-text {
        /* 縦書き設定 */
        writing-mode: vertical-rl;
        width: auto;
        height: auto;

        /* テキストのスタイル */
        color: #5C4033; /* 濃い茶色など、背景に合わせた色 */

        /* フォントサイズの調整（任意） */
        font-size: 18px;
        line-height: 2.2; /* 縦書きの行間 */
        padding: 20px;
        /* アルファベットを縦に表示 */
        text-orientation: upright;
    }

    .vertical-text h1 {
        font-size: 48px; /* 見出しのサイズ */
        margin-bottom: 30px; /* 見出しと本文の間隔 */
        font-weight: bold;
        text-shadow: 1px 1px 5px #ffffff;
    }

    .vertical-text p {
        font-size: 32px;
        font-weight: bold;
        margin: 0;
        text-shadow: 1px 1px 5px #ffffff;
    }

    .horizontal-text {
        display: flex;
        /* テキストが画面中央に配置されるように */
        align-items: center;
        justify-content: center;
        padding-bottom: 15vw;
    }

    .horizontal-text h1 {
        font-size: 48px; /* 見出しのサイズ */
        margin-bottom: 30px; /* 見出しと本文の間隔 */
        font-weight: bold;
        color: #5C4033; /* 濃い茶色など、背景に合わせた色 */
        text-shadow: 1px 1px 5px #ffffff;
}

</style>
