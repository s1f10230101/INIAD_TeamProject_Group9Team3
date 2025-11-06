<script>
import logo from "$lib/assets/icon4.png";
</script>

<header
  class="sticky top-0 left-0 flex justify-between items-center p-1 bg-[#fdfaf4]/80 border-b border-[#e0ddd7] w-full box-border sm:px-5"
>
  <a
    href="/"
    aria-label="トップページへ"
    class="flex items-center gap-3 text-[#3d3d3d] font-bold md:text-xs sm:text-base"
  >
    <img src={logo} alt="TRIP CANVAS ロゴ" class="h-10" />
    <span>TRIP CANVAS</span>
  </a>
  <nav>
    <ul class="flex gap-2 md:pb-2 md:gap-3">
      {@render a("/proposals", "旅行プラン生成")}
      {@render a("/facilities", "こだわりの施設一覧")}
      {@render a("/facilities/register", "新規施設登録")}
    </ul>
  </nav>
</header>

{#snippet a(/** @type {string} */ href, /** @type {string} */ text)}
  <a class=" text-[#5C4033] text-xs font-medium hover:text-[#967259]" {href}>
    <li>
      {text}
    </li>
  </a>
{/snippet}
