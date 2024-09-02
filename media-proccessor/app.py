from flask import Flask, request

app = Flask(__name__)

BASE_PATH = "/api/media-proccesor"

@app.route(f"{BASE_PATH}/health")
def hello_world():
    return {"status": "ok"}

@app.route(f"{BASE_PATH}/process-audio", methods=["POST"])
def process_audio():
    body = request.json
    file_id = body["file_id"]
    chat_id = body["chat_id"]
    return {"file_id": file_id, "chat_id": chat_id}

@app.route(f"{BASE_PATH}/process-picture", methods=["POST"])
def process_picture():
    body = request.json
    file_id = body["file_id"]
    chat_id = body["chat_id"]
    return {"file_id": file_id, "chat_id": chat_id}