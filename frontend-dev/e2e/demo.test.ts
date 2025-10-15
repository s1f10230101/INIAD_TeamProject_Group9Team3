import { expect, test } from '@playwright/test';

test('ルートページに0より多く、つまり1つ以上<h1>が存在することを確認する', async ({ page }) => {
	await page.goto('/');
	await expect(page.locator('h1').count()).toBeGreaterThan(0);
});
