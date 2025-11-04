<script>
import "../app.css";
import Header from "$lib/components/Header.svelte";
import Footer from "$lib/components/Footer.svelte";

let { children } = $props();
</script>

<div class="site-container">
  <Header />

  <main class="grow">
    {@render children()}
  </main>

  <Footer />
</div>

<svelte:head>
  <link rel="icon" href="/favicon.png" />
</svelte:head>

<style>
  .site-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh; /* 画面全体の高さを最低でも確保する */
  }
</style>
