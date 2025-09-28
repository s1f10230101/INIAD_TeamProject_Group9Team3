import os
import json
from typing import List
from fastapi import FastAPI
from fastapi.responses import StreamingResponse
from pydantic import BaseModel
from langchain_openai import ChatOpenAI, OpenAIEmbeddings
from langchain.prompts import PromptTemplate

# --- Configuration ---
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
OPENAI_API_BASE = os.getenv("OPENAI_API_BASE", "https://api.openai.com/v1")

# --- FastAPI App ---
app = FastAPI()

# --- Pydantic Models ---
class EmbedRequest(BaseModel):
    texts: List[str]

class GenerateRequest(BaseModel):
    question: str
    context: str

# --- LangChain Setup ---
embeddings = OpenAIEmbeddings(
    openai_api_key=OPENAI_API_KEY,
    base_url=OPENAI_API_BASE,
)

template = '''
以下の「検索結果」を参考にして、ユーザーの「要望」に合った旅行プランを提案してください。
提案は、具体的な場所やアクティビティを盛り込み、魅力的な文章で作成してください。

検索結果:
{context}

要望:
{question}
'''
prompt_template = PromptTemplate.from_template(template)

llm = ChatOpenAI(
    api_key=OPENAI_API_KEY,
    base_url=OPENAI_API_BASE,
    model="gpt-3.5-turbo",
    temperature=0.7,
    streaming=True,
)

# The chain is now simpler, taking context and question directly.
chain = prompt_template | llm

# --- API Endpoints ---

@app.post("/embed")
async def embed_texts(req: EmbedRequest):
    """
    Receives a list of texts and returns their embeddings.
    """
    try:
        vectors = embeddings.embed_documents(req.texts)
        return {"status": "success", "vectors": vectors}
    except Exception as e:
        return {"status": "error", "message": str(e)}, 500

async def stream_generator(question: str, context: str):
    """Generator that yields SSE-formatted streaming output from the LLM."""
    async for chunk in chain.astream({"question": question, "context": context}):
        content = chunk.content
        if content:
            # SSE (Server-Sent Events) format
            response_data = {"token": content}
            json_data = json.dumps(response_data, ensure_ascii=False)
            yield f"data: {json_data}\n\n"

@app.post("/generate-plan")
async def generate_plan(req: GenerateRequest):
    """Generates a travel plan based on user's prompt and provided context, and streams the response."""
    return StreamingResponse(stream_generator(req.question, req.context), media_type="text/event-stream")

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
