<script lang="ts">
import { enhance } from "$app/forms";
import Input from "$lib/components/Input.svelte";
import Textarea from "$lib/components/Textarea.svelte";
import type { PageProps } from "./$types";
let { form }: PageProps = $props();
</script>

<div class="rounded-2xl bg-white/90 p-4 space-y-3 w-full sm:w-5/6 lg:w-2/3 mt-6 mb-6 facility-register-page">

  <h1 class="text-center font-bold text-4xl text-[#5c4033]">施設情報登録</h1>
  <hr class="border-gray-300 border-t-2 my-4" />
  <form method="POST" use:enhance>
    <Input
      id="facilityName"
      label="施設名"
      name="facilityName"
      placeholder="施設名を入力"
      value={form?.facilityName?.toString() ?? ""}
      error={form?.missingName ? "名前は必須です" : ""}
      type="text"
    />
    <Input
      type="text"
      id="location"
      label="所在地"
      error=""
      name="location"
      placeholder="住所を入力"
      value={form?.location?.toString() ?? ""}
    />
    <Input
      type="text"
      id="hours"
      name="hours"
      placeholder="例: 9:00〜18:00"
      label="営業時間"
      error=""
      value={""}
    />

    <Input
      type="text"
      id="priceRange"
      name="priceRange"
      placeholder="例: ¥1,000〜¥3,000"
      label="価格帯"
      value={""}
      error={""}
    />

    <Textarea
      id="description"
      name="description"
      placeholder="施設の特徴や魅力を詳しく記入してください"
      value={form?.description?.toString() ?? ""}
      label="概要・説明"
      error={form?.missingDescription ? "記入して" : ""}
    />

    <div class="mt-2">
      <button
        type="submit"
        class="block w-full bg-[#6d4c41] rounded-4xl text-xl p-3 text-white hover:bg-[#5d4037] dark:text-gray-100" 
        >施設情報を登録する</button
      >
    </div>
  </form>
</div>
