import { expect, test } from "@playwright/test";

test("施設詳細ページにモックデータが表示されることを確認する", async ({
  page,
}) => {
  // MSWのモックデータにレビューが存在する東京タワーのID
  const facilityId = "c1b5c1c8-0b8f-4b1a-8b1a-0b8f4b1a8b1a";
  await page.goto(`/facilities/${facilityId}`);

  // 同様に「東京タワー」も表示されることを確認します
  const facility2 = page.getByText("東京タワー").first();
  await expect(facility2).toBeVisible();

  // 住所の一部が表示されているかも確認できます
  const address1 = page.getByText(/東京都港区芝公園４丁目２−８/);
  await expect(address1).toBeVisible();
});
