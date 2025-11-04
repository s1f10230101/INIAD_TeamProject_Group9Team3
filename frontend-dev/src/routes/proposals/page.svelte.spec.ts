import { render, screen, fireEvent, cleanup } from "@testing-library/svelte";
import { describe, it, expect, afterEach } from "vitest";
import ProposalsPage from "./+page.svelte";

describe("Proposals Page", () => {
  // 各テストの後にDOMをクリーンアップします
  afterEach(() => cleanup());

  // Test 1: 基本的な描画テスト
  it("ページが正しく描画され、初期の見出しが表示される", () => {
    render(ProposalsPage);
    const heading = screen.getByRole("heading", {
      name: "体験したい旅行体験をご自由にお書きください",
    });
    expect(heading).toBeInTheDocument();
  });

  // Test 2: フォーム送信とAPIレスポンスをテスト
  it("フォームを送信すると、生成された旅行プランがストリーミング表示される", async () => {
    render(ProposalsPage);

    // フォームの要素を取得します
    const textarea = screen.getByPlaceholderText("例：家族で温泉旅行");
    const submitButton = screen.getByRole("button", { name: "送信" });
    // ユーザーの入力をシミュレートします
    await fireEvent.input(textarea, {
      target: { value: "家族で楽しめる温泉旅行" },
    });

    // フォームの送信をシミュレートします
    await fireEvent.click(submitButton);

    // ボタンの表示が「生成中...」に変わることを確認します
    expect(
      screen.getByRole("button", { name: "生成中..." }),
    ).toBeInTheDocument();

    // ストリーミングされたレスポンスが画面に表示されるのを待ちます
    const responseText = await screen.findByText(/最高の東京旅行のプランです/);
    console.log(responseText);
    // レスポンスが表示されたことを確認します
    expect(responseText).toBeInTheDocument();

    // ボタンの表示が「送信」に戻っていることを非同期で待ち、確認します
    const finalButton = await screen.findByRole("button", { name: "送信" });
    expect(finalButton).toBeInTheDocument();
  });
});
