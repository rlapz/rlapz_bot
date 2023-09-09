import asyncio
import logging
import os
import sys
import dotenv

from aiohttp import web

from aiogram import Bot, Dispatcher
from aiogram.enums import ParseMode
from aiogram.webhook.aiohttp_server import (
    SimpleRequestHandler, setup_application
)

from handlers import anime_schedule
from handlers import general
import callback


dotenv.load_dotenv()
TOKEN = os.getenv("TOKEN")
SECRET = os.getenv("SECRET")
HOOK_URL = os.getenv("HOOK_URL")
HOOK_PATH = os.getenv("HOOK_PATH")
LISTEN_IP = os.getenv("LISTEN_IP")
LISTEN_PORT = int(os.getenv("LISTEN_PORT"))

HOOK = HOOK_URL + HOOK_PATH


async def startup_fn(bot: Bot) -> None:
    logging.info(f"setting up webhook: {HOOK}...")
    await bot.set_webhook(f"{HOOK}",
                          secret_token=SECRET,
                          drop_pending_updates=True)


def new_dispatcher() -> Dispatcher:
    ret = Dispatcher()

    # add handler(s)
    ret.include_routers(
        callback.r,
        general.r,
        anime_schedule.r,
    )

    return ret


def run_webhook() -> None:
    dp = new_dispatcher()
    dp.startup.register(startup_fn)

    bot = Bot(TOKEN, parse_mode=ParseMode.HTML)
    req = SimpleRequestHandler(dp, bot, secret_token=SECRET)

    app = web.Application()
    req.register(app, path=HOOK_PATH)

    setup_application(app, dp, bot=bot)
    web.run_app(app, host=LISTEN_IP, port=LISTEN_PORT)


async def run_polling() -> None:
    dp = new_dispatcher()
    bot = Bot(TOKEN, parse_mode=ParseMode.HTML)
    await dp.start_polling(bot)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)

    argv = sys.argv
    if len(argv) > 1 and argv[1] == "webhook":
        run_webhook()
    else:
        asyncio.run(run_polling())
