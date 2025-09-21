import os
import json
from fastapi import FastAPI
from fastapi.responses import StreamingResponse
from pydantic import BaseModel
from sqlalchemy import create_engine, text
from langchain_openai import ChatOpenAI, OpenAIEmbeddings
from langchain.schema.runnable import RunnableMap
from langchain.prompts import PromptTemplate
from langchain_community.vectorstores.pgvector import PGVector


# --- Configuration ---
DB_USER = os.getenv("POSTGRES_USER", "app_user")
DB_PASSWORD = os.getenv("POSTGRES_PASSWORD", "password")
DB_HOST = os.getenv("DB_HOST", "db")
DB_PORT = os.getenv("DB_PORT", "5432")
DB_NAME = os.getenv("POSTGRES_DB", "app_db")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
OPENAI_API_BASE = os.getenv("OPENAI_API_BASE", "https://api.openai.com/v1")

CONNECTION_STRING = f"postgresql+psycopg2://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}"
COLLECTION_NAME = "spots"

# --- FastAPI App ---
app = FastAPI()

# --- Pydantic Models ---
class PlanRequest(BaseModel):
    prompt: str

# --- LangChain Setup ---
embeddings = OpenAIEmbeddings(
    openai_api_key=OPENAI_API_KEY,
    base_url=OPENAI_API_BASE,
)
store = PGVector(
    connection_string=CONNECTION_STRING,
    embedding_function=embeddings,
    collection_name=COLLECTION_NAME,
)
retriever = store.as_retriever()

template = '''
以下の「検索結果」を参考にして、ユーザーの「要望」に合った旅行プランを提案してください。
提案は、具体的な場所やアクティビティを盛り込み、魅力的な文章で作成してください。

検索結果:
{context}

要望:
{question}
'''
prompt = PromptTemplate.from_template(template)
llm = ChatOpenAI(
    api_key=OPENAI_API_KEY,
    base_url=OPENAI_API_BASE,
    model="gpt-3.5-turbo",
    temperature=0.7,
    streaming=True,
)

chain = RunnableMap({
    "context": lambda x: retriever.get_relevant_documents(x["question"]),
    "question": lambda x: x["question"]
}) | prompt | llm

# --- API Endpoints ---
async def stream_generator(prompt: str):
    """LLMからのストリーミング出力をSSE形式でyieldするジェネレータ"""
    # astream_events coming soon to langchain, but for now, we use astream
    async for chunk in chain.astream({"question": prompt}):
        # chunk is an AIMessageChunk object
        content = chunk.content
        print(f"LLM chunk: {chunk}") # Debugging line
        print(content)
        if content:
            # SSE (Server-Sent Events) format
            yield f"data: {json.dumps({'token': content})}\\n\n"

@app.post("/generate-plan")
async def generate_plan(req: PlanRequest):
    """ユーザーの要望に基づいて旅行プランを生成し、ストリーミングで返す"""
    return StreamingResponse(stream_generator(req.prompt), media_type="text/event-stream")

@app.post("/private/update-embeddings")
async def update_embeddings():
    """データベース内のスポットの埋め込みを更新する（内部向け）"""
    try:
        engine = create_engine(CONNECTION_STRING)
        with engine.connect() as conn:
            results = conn.execute(text("SELECT id, name, description FROM spot")).fetchall()
            
            if not results:
                return {"status": "success", "message": "No spots found to update."}

            texts = [f"観光地名: {row[1]}\n説明: {row[2]}" for row in results]
            ids = [str(row[0]) for row in results]

            store.add_texts(texts=texts, ids=ids)

        return {"status": "success", "message": f"{len(texts)} embeddings updated."}
    except Exception as e:
        # In a real app, you'd want more robust error handling and logging
        return {"status": "error", "message": str(e)}, 500

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)