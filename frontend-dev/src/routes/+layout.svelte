<script>
	import '../app.css';
	import favicon from '$lib/assets/icon2.png';
	import Header from '$lib/components/Header.svelte';
	import Footer from '$lib/components/Footer.svelte';
	
	let { children } = $props();
</script>

<Header />

<main>

</main>

<Footer />

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children?.()}
