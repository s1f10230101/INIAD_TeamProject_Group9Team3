import { expect, test } from "@playwright/test";

test("施設一覧ページにモックデータが表示されることを確認する", async ({
  page,
}) => {
  await page.goto("/facilities");

  // MSWが返すモックデータの「東京タワー」が画面に表示されるのを待ちます
  const facility1 = page.getByText("東京タワー");
  await expect(facility1).toBeVisible();

  // 同様に「浅草寺」も表示されることを確認します
  const facility2 = page.getByText("浅草寺");
  await expect(facility2).toBeVisible();

  // 住所の一部が表示されているかも確認できます
  const address1 = page.getByText(/東京都港区芝公園/);
  await expect(address1).toBeVisible();
});
