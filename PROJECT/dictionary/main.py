import requests
from aiogram import Bot, types, executor
import os
import json
import subprocess
from aiogram.dispatcher import Dispatcher
from aiogram.dispatcher.filters import Text

bot_token = "6944612215:AAEkEraevZOqDpEV8-kqZMWthPKUkxPd1Tc"
bot = Bot(token=bot_token)
dp = Dispatcher(bot)


def dictionary(word):
    try:
        # Make the API request
        r = requests.get(f"https://api.dictionaryapi.dev/api/v2/entries/en/{word}")
        data = r.json()

        # Extracting relevant information
        word_info = data[0]  # Assuming there is only one result in the list

        word_text = f"Word: {word_info['word']}"

        # Extract and format phonetics information
        phonetics_info = word_info.get('phonetics', [])
        phonetics_text = "\nPhonetics:"
        for entry in phonetics_info:
            phonetics_text += f"\n- Text: {entry.get('text', '')}"
            phonetics_text += f"\n  Audio: {entry.get('audio', '')}"

        # Extract and format meanings information
        meanings_info = word_info.get('meanings', [])
        meanings_text = "\n\nMeanings:"
        for meaning_entry in meanings_info:
            part_of_speech = meaning_entry.get('partOfSpeech', '')
            meanings_text += f"\n{part_of_speech.capitalize()}:"

            definitions = meaning_entry.get('definitions', [])
            for definition in definitions:
                meanings_text += f"\n- Definition: {definition.get('definition', '')}"

        # Combine all information into a single text
        result_text = f"{word_text}{phonetics_text}{meanings_text}"

        return result_text

    except Exception as ex:
        print(ex)
        return "Not found in dictionary!"

@dp.message_handler(commands=['start'])
async def send_welcome(message: types.Message):
    chat_id = message.chat.id
    user_name = message.from_user.first_name
    await message.answer(f"Welcome {user_name}!\nBot designed to help with word recognition, detail extraction, and vocabulary input.\nCreated by @blvck_sudo"  )

@dp.message_handler()
async def echo_all(message: types.Message):
    # Foydalanuvchidan kelgan xabarni olamiz
    user_word = message.text
    answer=dictionary(user_word)
    await message.reply(answer)



if __name__ == '__main__':
    executor.start_polling(dp, skip_updates=True)

