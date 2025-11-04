import { expect, test } from "@playwright/test";

test("レビューページで施設詳細の表示とレビューの投稿ができる", async ({
  page,
}) => {
  // MSWのモックデータにレビューが存在する浅草寺のID
  const facilityId = "d2c6d2d9-1c9g-5c2b-9c2b-1c9g5c2b9c2b";
  await page.goto(`/facilities/reviews/${facilityId}`);

  // 1. データ表示のテスト
  // 施設名「浅草寺」が表示されることを確認
  await expect(page.getByText("浅草寺")).toBeVisible();

  // 詳細表示ボタンをクリック
  await page.getByRole("button", { name: /これ何のボタン？/ }).click();

  // 既存のレビューコメント「とても良かったです！」が表示されることを確認
  await expect(page.getByText("とても良かったです！")).toBeVisible();

  // 2. フォーム送信のテスト
  // フォーム要素を取得
  const contentTextarea = page.getByPlaceholder("詳細を入力");
  const ratingInput = page.getByLabel("施設評価");

  // フォームに入力
  await contentTextarea.fill("最高の体験でした！");
  await ratingInput.fill("4.8");

  // 「確認」ボタンをクリック
  await page.getByRole("button", { name: "確認" }).click();

  // 確認画面が表示されることを確認
  await expect(
    page.getByText("以下の内容で投稿します。よろしいですか？"),
  ).toBeVisible();
  await expect(page.getByText("最高の体験でした！")).toBeVisible();

  // 「投稿する」ボタンをクリック
  await page.getByRole("button", { name: "投稿する" }).click();

  // 投稿後に/facilitiesページへ遷移したことを確認
  await expect(page).toHaveURL("/facilities");
});
