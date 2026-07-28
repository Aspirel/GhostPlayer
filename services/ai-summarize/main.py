from fastapi import FastAPI
from pymongo import MongoClient
import os

app = FastAPI()

mongo_uri = os.getenv("MONGO_URI")
mongo = MongoClient(mongo_uri)
db = mongo["ghostplayer"]

@app.post("/summarize")
def summarize(payload: dict):
    # your logic here
    return {"status": "ok"}
