<script lang="ts">
interface Props {
  id: string;
  label: string;
  type: "text";
  value: string;
  name: string;
  error: string;
  placeholder: string;
}
let {
  id,
  label,
  type,
  value = $bindable(),
  name,
  error = "",
  placeholder = "",
}: Props = $props();
</script>

<div class="mb-6">
  <label for={id} class="block mb-2 font-bold text-[#5c4033] text-2xl leading-7">{label}</label>
  <input class="w-full p-3 rounded border box-border border-gray-300 text-base bg-[#fcfcfc] resize-y" {type} {id} {name} {placeholder} bind:value aria-invalid={!!error} />

  {#if error}
    <p class="text-red-600">{error}</p>
  {/if}
</div>
