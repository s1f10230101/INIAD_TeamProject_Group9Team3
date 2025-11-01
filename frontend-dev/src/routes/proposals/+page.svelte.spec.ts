import { render, screen, fireEvent } from '@testing-library/svelte';
import { describe, it, expect } from 'vitest';
import ProposalsPage from './+page.svelte';

describe('Proposals Page', () => {
    it('フォームを送信すると、生成された旅行プランが表示される', async () => {
    // ページコンポーネントを描画
    render(ProposalsPage);
    // フォームのテキストエリアと送信ボタンを取得
    const textarea = screen.getByLabelText('どんな旅行にしたいですか？');
    const submitButton = screen.getByRole('button', { name: '生成' });

    // テキストエリアに入力
    await fireEvent.input(textarea, { target: { value: '家族で楽しめる温泉旅行' } });
    // フォームを送信
    await fireEvent.click(submitButton);

    // MSWから返されたモックデータが画面に表示されるのを待機して確認
    // handlers.tsで定義した "最高の東京旅行" というテキストを探します
    const planTitle = await screen.findByText('最高の東京旅行');
    const planDescription = await screen.findByText(/東京の魅力を満喫する/);

    // テキストが存在することを確認
    expect(planTitle).toBeInTheDocument();
    expect(planDescription).toBeInTheDocument();
    });
});