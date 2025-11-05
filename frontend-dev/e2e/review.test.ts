import { expect, test } from "@playwright/test";

test("レビューページでレビューの表示ができる", async ({
  page,
}) => {
  // MSWのモックデータにレビューが存在する浅草寺のID
  const facilityId = "d2c6d2d9-1c9g-5c2b-9c2b-1c9g5c2b9c2b";
  await page.goto(`/facilities/${facilityId}/reviews`);

  // 1. データ表示のテスト
  // 既存のレビューコメント「とても良かったです！」が表示されることを確認
  await expect(page.getByText("とても良かったです！")).toBeVisible();

});
