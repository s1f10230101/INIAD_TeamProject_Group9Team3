<script>
import "../app.css";
import Header from "../lib/components/Header.svelte";
import Footer from "../lib/components/Footer.svelte";

let { children } = $props();
let backg = "../../static/back.png"
</script>

<div class="static-background min-h-screen bg-cover bg-fixed bg-center">
  <Header />

  <main class="container m-auto p-2 grow flex items-center flex-col">
    {@render children()}
  </main>

  <Footer />
</div>

<svelte:head>
  <link rel="icon" href="/favicon.png" />
</svelte:head>

<style>
  .static-background {
    background-image: url(backg); 
  }
</style>
cd 