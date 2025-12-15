import { expect, test } from "@playwright/test";

test("ページで詳細ページの表示ができる", async ({ page }) => {
  // MSWのモックデータにレビューが存在する浅草寺のID
  const facilityId = "d2c6d2d9-1c9g-5c2b-9c2b-1c9g5c2b9c2b";
  await page.goto(`/facilities/${facilityId}`);

  // loadingを待つ？謎バグ対処
  const load = page.getByTestId("loading");
  await expect(load).not.toBeVisible();

  const facility = page.getByText("浅草寺");
  await expect(facility).toBeVisible();

  const detail = page.getByText(/都内最古の寺院です/);
  await expect(detail).toBeVisible();
});
