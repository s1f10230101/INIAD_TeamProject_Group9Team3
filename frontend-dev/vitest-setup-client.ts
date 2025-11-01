/// <reference types="@vitest/browser/matchers" />
/// <reference types="@vitest/browser/providers/playwright" />

import "@testing-library/jest-dom/vitest";

import { worker } from "./src/mocks/browser";
import { beforeAll, afterEach, afterAll } from "vitest";

// すべてのテストの開始前にworkerを起動
beforeAll(() => worker.start());

// 各テストの終了後にハンドラーをリセット（テスト間の影響を防ぐため）
afterEach(() => worker.resetHandlers());

// すべてのテストの終了後にworkerを閉じる
afterAll(() => worker.stop());
