from aiogram import Bot, types, executor
import requests
from aiogram.dispatcher import Dispatcher
from aiogram.types import ParseMode
import re
import countryflag

bot_token = "6869839656:AAFGoYDjDtRzNnuHScRPTmyB3SvcSc-WP8U"
bot = Bot(token=bot_token)
dp = Dispatcher(bot)
name = "main"
usa = ["USA"]
usas = countryflag.getflag(usa)
rubli = ["Russia"]
rublis = countryflag.getflag(rubli)

def get_currency_rates():
    url = "https://cbu.uz/uz/arkhiv-kursov-valyut/json/"

    try:
        response = requests.get(url)
        response.raise_for_status()
        data = response.json()

        # Extracting currency rates
        dollar = extract_currency_data(data[0])
        euro = extract_currency_data(data[1])
        rubl = extract_currency_data(data[2])

        all_data = f"Valyuta: {dollar['name']}{usas}\nSana: {dollar['date']}\nQiymati: {dollar['rate']}\n\n" \
                   f"Valyuta: {euro['name']} €\nSana: {euro['date']}\nQiymati: {euro['rate']}\n\n" \
                   f"Valyuta: {rubl['name']}{rublis}\nSana: {rubl['date']}\nQiymati: {rubl['rate']}\n"

        return all_data, dollar['rate'], euro['rate'], rubl['rate']
    except requests.exceptions.RequestException as e:
        print(f"Xatolik yuz berdi: {e}")
        return None, None, None, None

def extract_currency_data(currency_data):
    return {
        'name': currency_data["CcyNm_UZ"],
        'date': currency_data["Date"],
        'rate': float(currency_data["Rate"])
    }

@dp.message_handler(commands=['start'])
async def send_welcome(message: types.Message):
    user_name = message.from_user.first_name
    await message.answer(f"Assalomualaykum {user_name}. /kurs kommandasi orqali valyuta kurslarni bilib oling \nCreated by @blvck_sudo")

@dp.message_handler(commands=['info', 'help'])
async def send_welcome(message: types.Message):
    await message.answer(f"Siz /kurs yordamida joriy valyutalar kursini korishingiz mumkun!\n Yoki siz usd,euro,rubl kurslarni sum ga hisoblashingiz mumkun!\nMisol uchun: dollar 10\nrubl 1000\neuro 5\nNamuna: dollar 10\n10 dollar 12..... sumga teng\n")

@dp.message_handler(commands=['kurs'])
async def send_welcome(message: types.Message):
    get_kurs, dollar_rate, euro_rate, rubl_rate = get_currency_rates()
    await message.answer(get_kurs)

def format_number(number):
    if isinstance(number, (int, float)):
        formatted_number = f"{number:.2f}"
        formatted_number = formatted_number.rstrip('0').rstrip('.')
    else:
        formatted_number = str(number)

    return formatted_number

async def process_currency_command(message: types.Message, currency_name, currency_rate):
    command_text = message.text
    match = re.search(r'\d+', command_text)

    if match:
        amount = int(match.group())
        result = amount * currency_rate
        formatted_amount = format_number(amount)
        formatted_result = format_number(result)
        response_text = f"{formatted_amount} {currency_name} {formatted_result} sum ga teng."
    else:
        response_text = "Xato! Siz noto'g'ri formatda yozdingiz."

    await message.reply(response_text, parse_mode=ParseMode.MARKDOWN)

@dp.message_handler(regexp=r'^dollar \d+$')
async def process_dollar_command(message: types.Message):
    _, dollar_rate, _, _ = get_currency_rates()
    await process_currency_command(message, 'dollar', dollar_rate)

@dp.message_handler(regexp=r'^euro \d+$')
async def process_euro_command(message: types.Message):
    _, _, euro_rate, _ = get_currency_rates()
    await process_currency_command(message, 'euro', euro_rate)

@dp.message_handler(regexp=r'^rubl \d+$')
async def process_rubl_command(message: types.Message):
    _, _, _, rubl_rate = get_currency_rates()
    await process_currency_command(message, 'rubl', rubl_rate)

if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)


