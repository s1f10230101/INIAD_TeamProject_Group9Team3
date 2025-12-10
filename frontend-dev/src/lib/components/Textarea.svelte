<script lang="ts">
interface Props {
  id: string;
  label: string;
  value: string;
  name: string;
  error: string;
  placeholder: string;
}
let {
  id,
  label,
  value = $bindable(),
  name,
  error = "",
  placeholder = "",
}: Props = $props();
</script>

<div class="mb-6">
  <label for={id} class="block mb-2 font-bold text-[#5c4033] text-2xl leading-7 dark:text-gray-100">{label}</label>
  <textarea class="w-full resize-none p-3 h-auto rounded-2xl border box-border border-gray-300 text-base bg-gray-50" {id} {name} {placeholder} bind:value aria-invalid={!!error} rows="15" >{value}</textarea>

  {#if error}
    <p class="text-red-600">{error}</p>
  {/if}
</div>
