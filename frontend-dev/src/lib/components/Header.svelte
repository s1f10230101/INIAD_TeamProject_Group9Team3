<script>
// ロゴ画像をインポートします。パスは実際の場所に合わせて調整してください。
import logo from "$lib/assets/icon4.png";
</script>

<header>
  <div class="logo-container">
    <a href="/" aria-label="トップページへ">
      <img src={logo} alt="TRIP CANVAS ロゴ" />
      <span>TRIP CANVAS</span>
    </a>
  </div>
  <nav>
    <ul>
      <li><a href="/proposals">旅行プラン生成</a></li>
      <li><a href="/facilities">こだわりの施設一覧</a></li>
      <li><a href="/auth/login">ログイン/新規登録</a></li>
    </ul>
  </nav>
</header>

<style>
  header {
    position: fixed; /* ヘッダーを画面に固定する */
    top: 0;          /* 画面の上端を基準にする */
    left: 0;         /* 画面の左端を基準にする */
    z-index: 100;
    display: flex;
    justify-content: space-between; /* 両端に寄せる */
    align-items: center; /* 上下中央揃え */
    padding: 10px 40px; /* 上下の余白、左右の余白 */
    background-color: #fdfaf4; /* 背景色を少しクリーム色に */
    border-bottom: 1px solid #e0ddd7; /* 下の境界線 */
    width: 100%;
    box-sizing: border-box; /* paddingを含めて幅を100%にする */
  }

  .logo-container a {
    display: flex;
    align-items: center;
    gap: 12px; /* アイコンと文字の間隔 */
    text-decoration: none;
    color: #3d3d3d;
    font-weight: bold;
    font-size: 20px;
  }

  .logo-container img {
    height: 40px; /* ロゴの高さを指定 */
    width: auto;
  }

  nav ul {
    list-style: none; /*箇条書きの点を消す*/
    margin: 0;
    padding: 0;
    display: flex;
    gap: 30px; /* メニュー項目の間隔 */
  }

  nav a {
    text-decoration: none;
    color: #5C4033; /* メインコンテンツと色を合わせる */
    font-size: 16px;
    font-weight: 500;
    transition: color 0.2s; /* ホバー時の色の変化を滑らかに */
  }

  nav a:hover {
    color: #967259; /* ホバー時の色 */
  }
</style>