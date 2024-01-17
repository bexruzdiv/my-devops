from aiogram import Bot, types, executor
import os
import json
import subprocess
from aiogram.dispatcher import Dispatcher
from aiogram.dispatcher.filters import Text
from db import insert_one_to_db


bot_token = "6439046083:AAGR-rFA6ZeQ0kT573mS-pYraSE_kvh1-j4"

bot = Bot(token=bot_token)
dp = Dispatcher(bot)

async def run_command(command):
    try:
        return subprocess.check_output(command, shell=True, text=True)
    except:
        return False
all_chat_ids=[]
@dp.message_handler(commands=['add'])
async def send_welcome(message: types.Message):
    chat_id = message.chat.id
    user_name = message.from_user.first_name
    is_bot = message.from_user.is_bot

    print(chat_id)

    if is_bot:
        print("I don't serve bots")
    else:
        chat_type = message.chat.type
        name = message.from_user.first_name

        if chat_type == "private":
            await message.answer(f"Hello {name}. Welcome!")
            data_to_insert = {
                "chat_id": chat_id,
                "user_name": user_name
            }

        
            insert_one_to_db("student", "student", data_to_insert)

        else:
            group_name = message.chat.title
            await message.reply(f"Hello from {name} to group {group_name}")

    await message.delete()




if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)