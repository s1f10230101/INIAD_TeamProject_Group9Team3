<script lang="ts">

	import { onMount } from 'svelte';

	import back10 from '$lib/assets/back10.png';
	import back11 from '$lib/assets/back11.png';
	import back12 from '$lib/assets/back12.png';
	import back13 from '$lib/assets/back13.png';

	import '../app.css';
	import favicon from '$lib/assets/icon4.png';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	
	let { children } = $props();

	const backgroundImages = [back10, back11, back12, back13];

  // 3. 状態変数を定義 
	let currentImage = $state(backgroundImages[0]);
  	let nextImage = $state(backgroundImages[1 % backgroundImages.length]);
  	let currentIndex = 0; // これはHTMLで使わないので `let` のままでOK
  	let isFadingOut = $state(false);

  // 4. onMountでタイマーをセット
  	onMount(() => {
    	const interval = setInterval(() => {
      		isFadingOut = true; 
      		setTimeout(() => {
        			currentIndex = (currentIndex + 1) % backgroundImages.length;
        			currentImage = backgroundImages[currentIndex];
        			nextImage = backgroundImages[(currentIndex + 1) % backgroundImages.length];
        			isFadingOut = false;
      		}, 1500); // CSSのアニメーション時間
    	}, 7000); // 切り替え間隔

    return () => {
      	clearInterval(interval);
    };
  });
</script>

<div class="background-wrapper">
  <div
    class="background-layer current-bg"
    style="background-image: url({currentImage})"
    class:fade-out={isFadingOut}
  ></div>
  <div
    class="background-layer next-bg"
    style="background-image: url({nextImage})"
  ></div>
</div>

<div class="site-container">
  	<Header />

  	<main>
    	{@render children()}
  	</main>

  	<Footer />
</div>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<style>
  /* 6. 背景フェード用のスタイルを追加 */
  .background-wrapper {
    position: fixed; /* 画面全体に固定 */
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
    z-index: -1; /* ★重要: すべてのコンテンツの背後に配置 */
  }

  .background-layer {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-size: cover;
    background-position: center center;
    opacity: 1;
    transition: opacity 1.5s ease-in-out; /* フェード時間 */
  }

  .current-bg.fade-out {
    opacity: 0;
  }

  .next-bg {
    opacity: 1;
  }

  /* 7. 既存のレイアウトスタイルを調整 */
  .site-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    /* 背景が透けて見えるように */
    background-color: transparent; 
  }

  main {
    flex-grow: 1;
    /* ヘッダー固定分の高さを確保 (以前の会話より) */
    padding-top: 61px; 
    /* 背景より手前に来るように */
    position: relative; 
    z-index: 1;
  }
</style>

