<script lang="ts">
  import { onMount } from "svelte";

  let theme: string|undefined = $state(undefined)
  onMount(() => theme = document.documentElement.dataset.theme)

  $effect(() =>
    {if(theme !== undefined) {
    document.documentElement.dataset.theme = theme
  }})
</script>

<footer
  class="sticky top-full box-border text-[0.7rem] flex justify-end p-2 z-2
bg-amber-50/80 text-primary-light-500
dark:bg-gray-800 dark:text-gray-100"
>
<!-- デバック用のボタンです。リロードするたびに設定は消えるので、ダークモードとライトモードの両方の見た目を確認したいとき以外使えません -->
<button onclick={() => theme = (theme==="light"? "dark":"light")}>
テーマ変更
</button>
  <p>
    © 2025 group9team3 &nbsp;&nbsp;&nbsp; 利用技術 ▶︎ Go, Svelte
  </p>
</footer>
