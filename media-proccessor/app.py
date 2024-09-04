from flask import Flask, request
from services.audio_service import AudioService

app = Flask(__name__)

BASE_PATH = "/api/media-proccesor"
audio_service = AudioService()

@app.route(f"{BASE_PATH}/health")
def hello_world():
    return {"status": "ok"}

@app.route(f"{BASE_PATH}/process-audio", methods=["POST"])
def process_audio():
    body = request.json
    file_path = body["file_path"]
    chat_id = body["chat_id"]
    transaction = audio_service.proccess(file_path)
    return transaction

@app.route(f"{BASE_PATH}/process-picture", methods=["POST"])
def process_picture():
    body = request.json
    file_path = body["file_path"]
    chat_id = body["chat_id"]
    return {"file_path": file_path, "chat_id": chat_id}