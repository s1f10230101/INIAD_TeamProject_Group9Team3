<script>
import "../app.css";
import Header from "./Header.svelte";
import Footer from "./Footer.svelte";

let { children } = $props();
</script>

<!-- この下にfont-robotoを追加すると全体がrobotoフォントになるよ -->
<div class="min-h-screen bg-[url('/back.png')] bg-cover bg-fixed bg-center">
  <Header />

  <main class="container m-auto p-2 grow flex items-center flex-col">
    {@render children()}
  </main>

  <Footer />
</div>

<svelte:head>
  <link rel="icon" href="/favicon.png" />
</svelte:head>
