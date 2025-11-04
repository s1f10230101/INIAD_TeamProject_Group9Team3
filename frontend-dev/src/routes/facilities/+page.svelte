<script lang="ts">
import backgroundImage from "$lib/assets/back10.png";
import FacilityList from "./FacilityList.svelte";

let { data } = $props();
const { facilitiesWithRatings } = data;
</script>

<div
  class="full-screen-background"
  style="--background-url: url('{backgroundImage}')"
>
  <main class="center-content">
    <div class="review-main-container">
      {#await facilitiesWithRatings}
        <p>loading</p>
      {:then facilities}
        <FacilityList facilities={facilities.filter((e)=>e !==undefined)} />
      {:catch err}
        {err}
      {/await}
    </div>
  </main>
</div>

<style>
  .full-screen-background {
    background-size: cover;
    background-position: center center;
    background-attachment: fixed;
    background-image: var(--background-url);
    padding-top: 5vw;
    display: flex;
    justify-content: center; /* 横方向の中央寄せ */
    align-items: center; /* 縦方向の中央寄せ */
    padding-bottom: 5vw;
  }

  .center-content {
    padding: 5vw;
  }

  .review-main-container {
    width: 60vw; /* フォーム全体の幅 */
    padding: 1.5vh 1vw;
    background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
    border-radius: 10px; /* 角丸 */
    box-shadow: 0 4px 15px #5c4033; /* 軽い影 */
    margin-bottom: 30px;
  }

</style>
