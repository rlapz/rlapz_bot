from aiogram import Bot, Router
from aiogram.types import CallbackQuery

import utils

from handlers import anime_schedule
from handlers import general

r = Router()


@r.callback_query()
async def callback_fn(cb: CallbackQuery, bot: Bot) -> None:
    if cb.data is None:
        await cb.answer(f"no data")
        return

    # code:udata
    id, udata = utils.tok_1(cb.data, ":")

    if id == anime_schedule.CALLBACK_ID:
        await anime_schedule.run_cb(bot, cb.message, udata)
    elif id == general.CALLBACK_ID:
        await general.run_cb(bot, cb.message, udata)
    else:
        await cb.answer(f"invalid callback id: {id}")
