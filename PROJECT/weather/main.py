import requests
import datetime
from config import open_weather_token
from aiogram import Bot, types, executor
import os
import json
import subprocess
from aiogram.dispatcher import Dispatcher
from aiogram.dispatcher.filters import Text
bot_token = "6426148954:AAEIJTT0YoGRkjJqYBMuagY-xScMX5eXxo4"
bot = Bot(token=bot_token)
dp = Dispatcher(bot)
def get_weather(city, open_weather_token):
    try:
        r = requests.get(
            f"https://api.openweathermap.org/data/2.5/weather?q={city}&appid={open_weather_token}&units=metric"
        )
        data =r.json()
        city = data["name"]
        cur_weather=data["main"]["temp"]
        sunrise_time=datetime.datetime.fromtimestamp(data["sys"]["sunrise"])
        sunset_time=datetime.datetime.fromtimestamp(data["sys"]["sunset"])
        humidity=data["main"]["humidity"]
        pressure=data["main"]["pressure"]
        country=data["sys"]["country"]
        all_in=(f"Shahar: {city}\n Davlat: {country}\n Harorat: {cur_weather}\n Namlik: {humidity}\n Bosim: {pressure}\n Quyosh chiqishi: {sunrise_time}\n Quyosh botishi: {sunset_time}")
        return all_in;
    except Exception as ex:
        print(ex)
        print("Wrong city! Check your location")

@dp.message_handler(commands=['start'])
async def send_welcome(message: types.Message):
    chat_id = message.chat.id
    user_name = message.from_user.first_name
    await message.answer(f"Assalomualaykum {user_name}. Hohlagan shahar yoki davlat nomini kiriting! \nCreated by @blvck_sudo"  )
@dp.message_handler()
async def echo_all(message: types.Message):
    # Foydalanuvchidan kelgan xabarni olamiz
    cites = message.text
    answer=get_weather(cites, open_weather_token)
    await message.reply(answer)


if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)









