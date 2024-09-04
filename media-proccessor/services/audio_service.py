import json
import os
import requests
import whisper
from whisper import Whisper
from openai import OpenAI
from dotenv import load_dotenv

load_dotenv()

client = OpenAI()
class AudioService():
    model:Whisper
    audio_media_dir:str = "audio_media_files"
    def __init__(self):
        print(os.environ.get("OPENAI_API_KEY"))
        self.model = whisper.load_model("base")

    def proccess(self, file_path:str) -> str:
        response = requests.get(file_path, stream=True)

        file_name = file_path.split("/")[-1]
        local_path = f"{self.audio_media_dir}/{file_name}"

        with open(local_path, "wb") as file:
            file.write(response.content)
        result = self.model.transcribe(local_path)
        transaction = self.infer_transaction(result.get("text", "No text found"))
        return transaction
    
    def infer_transaction(self, text:str) -> dict:
        message = "Based on the text, generate a JSON response with the fields 'amount', 'category' , 'type', type must be 'expense' or 'income'. \n This is the message: " + text
        completion = client.chat.completions.create(
        model="gpt-3.5-turbo",
        temperature=0,
        messages=[
            {
            "role": "user",
            "content": message
            }
        ]
        )
        res = json.loads(completion.choices[0].message.content)
        return res