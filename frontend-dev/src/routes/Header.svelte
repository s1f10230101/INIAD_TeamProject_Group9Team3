<script>

// 🍔 モバイルメニューの開閉状態を管理する変数
let menuOpen = $state(false);

// 画面サイズが変わったときにメニューを閉じる処理（Svelte 5のイベントハンドリング）
// window:resizeイベントはSvelteKitのSSRと互換性を持たせるため、onMountで囲む方が安全ですが、
// $effectを使用するとSvelte 5の記法としてより簡潔になります。
// ここでは簡易的にmenuOpenを切り替える関数を定義します。
function toggleMenu() {
  menuOpen = !menuOpen;
}

function closeMenu() {
  menuOpen = false;
}
</script>

<header
  class="sticky top-0 left-0 z-50 flex flex-col sm:flex-row justify-between items-center p-2
  bg-amber-50/90 border-b border-amber-100 w-full box-border shadow-md sm:px-5
  dark:bg-gray-800/90 dark:border-gray-700"
>
  <div class="flex justify-between items-center w-full sm:w-auto">
    <a
      href="/"
      aria-label="トップページへ"
      class="flex items-center gap-3 font-bold text-base"
      onclick={closeMenu}
    >
      <img src="/favicon.png" alt="TRIP CANVAS ロゴ" class="h-10" />
      <span>TRIP CANVAS</span>
    </a>

    <button
      onclick={toggleMenu}
      class="sm:hidden p-2 hover:text-primary-light-700 focus:outline-none"
      aria-expanded={menuOpen}
      aria-controls="navigation-menu"
      aria-label="メニューを開閉"
    >
      <svg
        class="w-6 h-6"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        {#if menuOpen}
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          />
        {:else}
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h16"
          />
        {/if}
      </svg>
    </button>
  </div>

  <nav
    id="navigation-menu"
    class:hidden={!menuOpen}
    class="w-full sm:w-auto mt-2 sm:mt-0
           sm:flex sm:items-center sm:visible sm:h-auto
           transition-all duration-300 ease-in-out"
  >
    <ul

      class="flex flex-col sm:flex-row gap-2 sm:gap-3 p-2 sm:p-0
             bg-amber-50 sm:bg-transparent border border-amber-100 sm:border-none rounded-md dark:bg-gray-800/90 dark:border-gray-700"
    >
      {@render a("/proposals", "旅行プラン生成", closeMenu)}
      {@render a("/facilities", "こだわりの施設一覧", closeMenu)}
      {@render a("/facilities/register", "新規施設登録", closeMenu)}
    </ul>
  </nav>
</header>

{#snippet a(/** @type {string} */ href, /** @type {string} */ text, /** @type {() => void} */ on_click)}
  <a
    class="text-sm font-medium hover:text-primary-light-700
           py-2 px-3 sm:py-0 sm:px-0 block sm:inline-block
           rounded-md transition-colors duration-200"
    {href}
    onclick={on_click}
  >
    <li>
      {text}
    </li>
  </a>
{/snippet}
