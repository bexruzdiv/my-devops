from flask import Flask, request, jsonify
from flask_cors import CORS
from flask import render_template, jsonify
app = Flask(__name__)
CORS(app)
import os
import json
import subprocess
from db import get_all_chat_ids
import requests
bot_token = "6439046083:AAGR-rFA6ZeQ0kT573mS-pYraSE_kvh1-j4"
base_url = f"https://api.telegram.org/bot{bot_token}"

def send_message(chat_id, text_to_send):
    url = f"{base_url}/sendMessage"
    params = {
        'chat_id': chat_id,
        'text': text_to_send,
    }

    try:
        response = requests.post(url, params=params)
        response.raise_for_status()
        print(f"Message sent to chat ID: {chat_id}")
    except requests.exceptions.RequestException as e:
        print(f"Error sending message to chat ID {chat_id}: {e}")

def send_message_to_all_chat_ids(text_to_send):
    all_chat_ids = get_all_chat_ids("student", "student")
    print(all_chat_ids)
    for target_chat_id in all_chat_ids:
        send_message(target_chat_id, text_to_send)




@app.route('/send', methods=['POST'])
def delete_student():
    data = request.json

    first_name = data['first_name']

    send_message_to_all_chat_ids(first_name)

    return jsonify({"message": "Record  successfully"})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)




