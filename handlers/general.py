from aiogram import Bot, Router, filters
from aiogram.types import Message
from aiogram.utils import markdown as md
from aiogram.types import Message, InlineKeyboardMarkup
from aiogram.utils.keyboard import InlineKeyboardBuilder

import utils

CALLBACK_ID = "1"

#
# Routes
#
r = Router()


#
# Messages
#
@r.message(filters.CommandStart())
async def cmd_start_fn(msg: Message) -> None:
    await msg.answer("Bacott :)")


def gen_kbd(id: str, udata: str) -> InlineKeyboardMarkup:
    builder = InlineKeyboardBuilder()
    builder.button(
        text="Press",
        callback_data=f"{id}:{udata}"
    )

    return builder.as_markup()


@r.message(filters.Command("test_btn"))
async def cmd_test_btn(msg: Message) -> None:
    markup = gen_kbd(CALLBACK_ID, f"@{msg.from_user.username}")
    await msg.reply("...", reply_markup=markup)


async def run_cb(bot: Bot, msg: Message, udata: str) -> None:
    await bot.edit_message_text(chat_id=msg.chat.id,
                                message_id=msg.message_id,
                                text=udata)
