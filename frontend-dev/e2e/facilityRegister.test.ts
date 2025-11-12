import { test } from "@playwright/test";

test("施設登録ページにフォームが表示されることを確認する", async ({ page }) => {
  await page.goto(`/facilities/register`);

  const nameInput = page.locator("#facilityName");
  const locateInput = page.locator("#location");
  const hours = page.locator("#hours");
  const range = page.locator("#priceRange");
  const description = page.locator("#description");
  const submitButton = page.locator('button[type="submit"]');

  await nameInput.fill("a");
  await locateInput.fill("a");
  await hours.fill("a");
  await range.fill("a");
  await description.fill("a");

  await submitButton.click();
});
