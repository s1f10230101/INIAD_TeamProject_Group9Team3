import { page } from "@vitest/browser/context";
import { describe, expect, it } from "vitest";
import { render } from "vitest-browser-svelte";
import Page from "./+page.svelte";

describe("/+page.svelte", () => {
  it("h1タグが少なくとも1つあることをテスト", async () => {
    render(Page);

    const heading = page.getByRole("heading", { level: 1 }).first();
    await expect.element(heading).toBeInTheDocument();
  });
});
