<script>
import backgroundImage from "$lib/assets/back10.png";
import client from '$lib/api/client';
    import { preventDefault } from "svelte/legacy";

let prompt = '';
let aiResponse = '';
let isLoading = false;

async function handleSubmit() {
    if (!prompt) return;
    isLoading = true;
    aiResponse = '';

    const { response } = await client.POST('/plans', {
        body: {
            prompt:prompt
        },
        parseAs: 'stream'
    });

    if (response.body) {
        const reader = response.body.getReader();
        const decoder = new TextDecoder();

        try {
            while (true) {
                const {done, value} = await reader.read();
                if (done) break;

                const chunk = decoder.decode(value, { stream: true });
                const lines = chunk.split('\n\n');

                for (const line of lines) {
                    if (line.startsWith('data:')) {
                        const data = line.substring(5).trim();
                        try {
                            const parsed = JSON.parse(data);
                            if (parsed.text) {
                                aiResponse += parsed.text;
                            }
                        } catch (e) {
                            console.error('SSE データの解析エラー:', e);
                        }
                    } else if (line.includes('event: done')) {
                        reader.cancel();
                        return;
                    }
                }
            }
        } catch (error) {
            console.error('ストリーム読み込みエラー:', error);
        } finally {
            isLoading = false;
        }
    } else {
        isLoading = false;
    }
}
</script>

<div class="full-screen-background" style="--background-url: url('{backgroundImage}')" >
    <main class="center-content">
        <h1>体験したい旅行体験をご自由にお書きください</h1>
        <!--<input type="text" id="xx" class="form" name="form" placeholder="例:  家族で温泉旅行" required>-->
        <form on:submit|preventDefault={handleSubmit}>
            <input 
                type="text"
                bind:value={prompt}
                class="form"
                placeholder="例：家族で温泉旅行"
            />
            <button type="submit" class="submit-button" disabled={isLoading}>
                {isLoading ? '生成中...' : '送信'}
            </button>
        </form>
        <div class="ai-response-container">
            <div class="ai-response-box">
                {aiResponse}
            </div>
        </div>
    </main>
</div>

<style>
    .full-screen-background{
        background-size: cover;
        background-position: center center; 
        background-attachment: fixed;
        background-image: var(--background-url); 
        padding-top: 10vw;
        display: flex;
        justify-content: center; /* 横方向の中央寄せ */
        align-items: center;   /* 縦方向の中央寄せ */
    }
    
    .center-content {
        padding: 5vw;
    }

    .center-content h1 {
        font-weight: bold;
        color: #5C4033;
        font-size: 28px;
        display: flex;
        justify-content: center; /* 横方向の中央寄せ */
        padding-bottom: 1vw;
        text-shadow: 1px 1px 5px #ffffff;
    }

    .form {
        width: 60vw; /* フォーム全体の幅 */
        padding: 1.5vh 1vw;
        background-color: rgba(255, 255, 255, 0.9); /* 半透明の白背景 */
        border-radius: 10px; /* 角丸 */
        box-shadow: 0 4px 15px #5C4033; /* 軽い影 */
        margin-bottom: 30px;
    }

    .ai-response-container {
        width: 60vw;
        margin-top: 20px;
    }


	.ai-response-box {
		width: 100%;
		min-height: 20vh; /* 最低の高さを確保 */
		padding: 20px;
		background-color: rgba(255, 255, 255, 0.95); /* より不透明な白背景 */
		border-radius: 10px;
		box-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
		font-size: 16px;
		line-height: 1.6;
		color: #5C4033;
        white-space: pre-wrap; /* テキストの改行を有効にする */
        padding-bottom: 20vw;
	}





</style>