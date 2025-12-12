<script lang="ts">
import "../app.css";
import Header from "./Header.svelte";
import Footer from "./Footer.svelte";

let { children } = $props();
</script>
<div
  class="min-h-screen bg-[url('/back.png')] bg-cover bg-fixed bg-center dark:bg-[url('/back-dark.png')]
  text-primary-light-500 dark:text-gray-100 dark:bg-black"
>
  <Header />
  <main class="container m-auto grow flex items-center flex-col pb-10 px-2">
    {@render children()}
  </main>

  <Footer />
</div>

<svelte:head>
  <link rel="icon" href="/favicon.png" />
</svelte:head>
