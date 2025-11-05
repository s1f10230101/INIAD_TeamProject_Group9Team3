import { expect, test } from "@playwright/test";

test("施設詳細ページにモックデータが表示されることを確認する", async ({
  page,
}) => {
  // MSWのモックデータにレビューが存在する浅草寺のID
  const facilityId = "d2c6d2d9-1c9g-5c2b-9c2b-1c9g5c2b9c2b";
  await page.goto(`/facilities/${facilityId}`);


  // 同様に「浅草寺」も表示されることを確認します
  const facility2 = page.getByText("浅草寺");
  await expect(facility2).toBeVisible();

  // 住所の一部が表示されているかも確認できます
  const address1 = page.getByText(/東京都台東区浅草２丁目３−１/);
  await expect(address1).toBeVisible();
});
