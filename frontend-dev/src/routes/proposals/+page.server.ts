import client, { streamingRecvHelper } from "$lib/api/client";
import { fail } from "@sveltejs/kit";

export const actions = {
  default: async ({ request }) => {
    const recvData = await request.formData();
    const promptData = recvData.get("prompt");
    if (!promptData || promptData.toString()==="") {
      return fail(400, {
        invaild: "ひっす"
      })
    };
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
