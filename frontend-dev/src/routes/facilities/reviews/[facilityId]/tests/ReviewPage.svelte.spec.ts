import {
    render,
    screen,
    fireEvent,
    cleanup,
    waitFor,
} from "@testing-library/svelte";
import { describe, it, expect, afterEach, vi } from "vitest";
import ReviewPage from "../+page.svelte";

// SvelteKitの$pageストアをモックします
// readableをファクトリ関数の中で動的にインポートして、巻き上げ問題を回避します
vi.mock("$app/stores", async (importOriginal) => {
    const { readable } = await import("svelte/store");
    const original = await importOriginal();
    return {
        original,
        page: readable({
            params: { facilityId: "c1b5c1c8-0b8f-4b1a-8b1a-0b8f4b1a8b1a" }, // テスト用のID
        }),
    };
});

// SvelteKitのナビゲーション機能をモックします
vi.mock("$app/navigation", () => ({
    goto: vi.fn(),
}));

// モックされた関数を後からインポートします
import { goto } from "$app/navigation";

describe("Review Submission Page", () => {
    // 各テストの後にDOMとモックをクリーンアップします
    afterEach(() => {
        cleanup();
        vi.clearAllMocks();
    });

    it("マウント時に施設詳細と既存レビューが表示される", async () => {
        render(ReviewPage);

        // APIから取得した施設名が表示されるのを待ちます
        const facilityName = await screen.findByText("東京タワー");
        expect(facilityName).toBeInTheDocument();

        // 詳細表示ボタンをクリックします
        const detailsButton = screen.getByRole("button", {
            name: /東京タワー/,
        });
        await fireEvent.click(detailsButton);

        // APIから取得したレビューコメントが表示されるのを待ちます
        const reviewComment = await screen.findByText("とても良かったです！");
        expect(reviewComment).toBeInTheDocument();
    });

    it("フォームを入力して確認し、レビューを投稿すると画面遷移が起こる", async () => {
        // window.alertをスパイし、何もしないようにします
        vi.spyOn(window, "alert").mockImplementation(() => {});

        render(ReviewPage);

        // ページがAPIデータを読み込むのを待ちます
        await screen.findByText("東京タワー");

        // フォームの要素を取得します
        const contentTextarea = screen.getByPlaceholderText("詳細を入力");
        const ratingInput = screen.getByLabelText("施設評価");

        // ユーザー入力をシミュレートします
        await fireEvent.input(contentTextarea, {
            target: { value: "最高の体験でした！" },
        });
        await fireEvent.input(ratingInput, { target: { value: "4.8" } });

        // 「確認」ボタンをクリックします
        const confirmButton = screen.getByRole("button", { name: "確認" });
        await fireEvent.click(confirmButton);

        // 確認画面が表示されたことを確認します
        expect(
            screen.getByText("以下の内容で投稿します。よろしいですか？"),
        ).toBeInTheDocument();
        expect(screen.getByText("最高の体験でした！")).toBeInTheDocument();

        // 「投稿する」ボタンをクリックします
        const postButton = screen.getByRole("button", { name: "投稿する" });
        await fireEvent.click(postButton);

        // `goto`が呼ばれるまで待機して、確認します
        await waitFor(() => {
            expect(goto).toHaveBeenCalledWith("/facilities");
        });
    });
});
