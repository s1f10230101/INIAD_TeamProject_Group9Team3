import { expect, test } from "@playwright/test";

test("プロポーザルページでフォームを送信し、AIの応答が表示されることを確認する", async ({
  page,
}) => {
  await page.goto("/proposals");

  // フォームの要素を取得します
  const textarea = page.getByPlaceholder("例：家族で温泉旅行");
  const submitButton = page.getByRole("button", { name: "送信" });

  // ユーザーの入力をシミュレートします
  await textarea.fill("家族で楽しめる温泉旅行");

  // フォームの送信をシミュレートします
  await submitButton.click();

  // ストリーミングされたレスポンスが画面に表示されるのを待ちます
  // MSWのハンドラで設定されているモックデータの一部を期待します
  const responseText = page.getByText(/最高の東京旅行のプランです/);
  await expect(responseText).toBeVisible();
});
