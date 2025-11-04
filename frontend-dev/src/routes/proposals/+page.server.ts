import client, { streamingRecvHelper } from "$lib/api/client";

export const actions = {
  default: async ({ request }) => {
    const recvData = await request.formData();
    const promptData = recvData.get("prompt");
    if (!promptData) return;
    const prompt = promptData.toString();
    const { response } = await client.POST("/plans", {
      body: {
        prompt: prompt,
      },
      parseAs: "stream",
    });

    let acc = "";
    await streamingRecvHelper(response, (sseText) => {
      acc += sseText;
    });

    return { res: acc };
  },
};
