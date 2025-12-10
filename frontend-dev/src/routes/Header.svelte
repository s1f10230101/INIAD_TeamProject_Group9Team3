<script>
import logo from "$lib/assets/icon4.png";

let menuOpen = false;
function toggleMenu() {
  menuOpen = !menuOpen;
}

function closeMenu() {
  menuOpen = false;
}
</script>

<header
  class="sticky top-0 left-0 z-50 flex flex-col sm:flex-row justify-between items-center p-2 bg-[#fdfaf4]/90 border-b border-[#e0ddd7] w-full box-border shadow-md sm:px-5"
>
  <div class="flex justify-between items-center w-full sm:w-auto">
    <a
      href="/"
      aria-label="トップページへ"
      class="flex items-center gap-3 text-[#3d3d3d] font-bold text-base"
      on:click={closeMenu}
    >
      <img src={logo} alt="TRIP CANVAS ロゴ" class="h-10" />
      <span>TRIP CANVAS</span>
    </a>

    <button
      on:click={toggleMenu}
      class="sm:hidden p-2 text-[#5C4033] hover:text-[#967259] focus:outline-none"
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
             bg-[#fdfaf4] sm:bg-transparent border border-[#e0ddd7] sm:border-none rounded-md"
    >
      {@render a("/proposals", "旅行プラン生成", closeMenu)}
      {@render a("/facilities", "こだわりの施設一覧", closeMenu)}
      {@render a("/facilities/register", "新規施設登録", closeMenu)}
    </ul>
  </nav>
</header>

{#snippet a(/** @type {string} */ href, /** @type {string} */ text, /** @type {() => void} */ on_click)}
  <a
    class="text-[#5C4033] text-sm font-medium hover:text-[#967259] 
           py-2 px-3 sm:py-0 sm:px-0 block sm:inline-block 
           rounded-md transition-colors duration-200"
    {href}
    on:click={on_click}
  >
    <li>
      {text}
    </li>
  </a>
{/snippet}