<script>
  import "../app.css";
  import favicon from "$lib/assets/icon4.png";
  import Header from "$lib/components/Header.svelte";
  import Footer from "$lib/components/Footer.svelte";

  let { children } = $props();
</script>

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
  .site-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh; /* 画面全体の高さを最低でも確保する */
  }

  main {
    flex-grow: 1; /* main要素が残りの高さをすべて埋めるように伸縮する */
  }
</style>
