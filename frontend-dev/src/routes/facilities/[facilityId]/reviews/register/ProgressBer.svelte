<div class="progress-bar-container w-full">
  <div
    class="progress-step"
    class:is-active={true}
  >
    1.情報入力
  </div>
  <div class="progress-step" class:is-active={false}>2.確認</div>
  <div class="progress-step">3.完了</div>
</div>

<style>
  .progress-bar-container {
    display: flex;
    justify-content: center; /* フォームに揃えて中央に配置 */
    margin-bottom: 20px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
  }

  .progress-step {
    flex-grow: 1;
    padding: 10px 30px 10px 20px;
    font-size: 16px;
    font-weight: bold;
    color: #5c4033; /* デフォルトの文字色 */
    background-color: #e0e0e0; /* デフォルトの背景色（灰色） */
    text-align: center;
    position: relative;
    /* 矢印の形状を作るために必須 */
    clip-path: polygon(
      0 0,
      calc(100% - 15px) 0,
      100% 50%,
      calc(100% - 15px) 100%,
      0 100%,
      15px 50%
    );
    z-index: 10;
  }

  /* 1番目のステップは左端を四角にする */
  .progress-bar-container .progress-step:first-child {
    padding-left: 15px;
    clip-path: polygon(
      0 0,
      calc(100% - 15px) 0,
      100% 50%,
      calc(100% - 15px) 100%,
      0 100%
    );
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
    background-color: #e0e0e0;
    color: #5c4033;
    /* ... 他のスタイル ... */
  }


  /* 現在アクティブなステップのスタイル */
  .progress-step.is-active {
    background-color: #6d4c41; /* 現在地を示す濃い茶色 */
    color: white;
  }
</style>
