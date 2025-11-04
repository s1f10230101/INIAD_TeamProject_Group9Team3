import { render, screen, cleanup } from "@testing-library/svelte";
import { describe, it, expect, afterEach } from "vitest";
import FacilitiesPage from "./+page.svelte";

describe("Facilities Page", () => {
  // 各テストの後にDOMをクリーンアップします
  afterEach(() => cleanup());

  it("マウント時に施設の一覧を取得して表示する", async () => {
    render(FacilitiesPage);

    // MSWが返すモックデータの「東京タワー」が画面に表示されるのを待ちます
    const facility1 = await screen.findByText("東京タワー");
    expect(facility1).toBeInTheDocument();

    // 同様に「浅草寺」も表示されることを確認します
    const facility2 = await screen.findByText("浅草寺");
    expect(facility2).toBeInTheDocument();

    // 住所の一部が表示されているかも確認できます
    const address1 = await screen.findByText(/東京都港区芝公園/);
    expect(address1).toBeInTheDocument();
  });
});
