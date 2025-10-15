<script>
  import backgroundImage from '$lib/assets/back10.png'; 
</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main>

    </main>
</div>

<style>
    .full-screen-background{
        background-size: cover;
        background-position: center center; 
        background-attachment: fixed;
        background-image: var(--background-url); 
    }


</style>