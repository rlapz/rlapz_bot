from aiogram import Bot, Router, filters
from aiogram.types import Message, InlineKeyboardMarkup
from aiogram.utils.keyboard import InlineKeyboardBuilder

import utils
from utils import err

TARGET_URL = "https://api.jikan.moe/v4/schedules"
TARGET_LIMIT = 3
CALLBACK_ID = "2"

r = Router()


class AnimeSched:
    day: str
    page: int
    has_next: bool
    result: str

    def __init__(self, day: str = "", page: int = -1,
                 has_next: bool = False, result: str = "") -> None:
        self.day = day
        self.page = page
        self.has_next = has_next
        self.result = result


async def __get_schedule(day: str = "", page: int = 1) -> AnimeSched:
    if len(day) == 0:
        day = utils.get_current_day()

    url = f"{TARGET_URL}?page={page}&filter={day}&limit={TARGET_LIMIT}&kids=false"
    resp, body = await utils.http_request_json(url)
    if resp.status != 200:
        raise f"invalid response status: {resp.status}"

    pagination = body['pagination']
    curr_page = int(pagination['current_page'])
    has_next = bool(pagination['has_next_page'])
    items = pagination["items"]

    res = f"<pre>Schedule: {day.title()}</pre>\n"
    res += f"<pre>Page    : {curr_page}</pre>\n"
    res += f"<pre>Limit   : {items['per_page']}</pre>\n"
    res += f"<pre>Total   : {items['total']}</pre>\n\n"

    for i, d in enumerate(body['data']):
        res += f"<b>{i + 1}. <a href='{d['url']}'>{d['title']}</a></b>\n"
        res += f"<pre>  Japanese : {d['title_japanese']}</pre>\n"
        res += f"<pre>  Type     : {d['type']}</pre>\n"
        res += f"<pre>  Episodes : {d['episodes']}</pre>\n"
        res += f"<pre>  Source   : {d['source']}</pre>\n"
        res += f"<pre>  Duration : {d['duration']}</pre>\n"
        res += f"<pre>  Score    : {d['score']}</pre>\n"
        res += f"<pre>  Status   : {d['status']}</pre>\n"
        res += f"<pre>  Rating   : {d['rating']}</pre>\n"

        genres = utils.join_list_dict(d['genres'], 'name')
        res += f"<pre>  Genre    : {genres}</pre>"

        themes = utils.join_list_dict(d['themes'], 'name')
        res += f"<pre>  Themes   : {themes}</pre>"

        dgraphics = utils.join_list_dict(d['demographics'], 'name')
        res += f"<pre>  Dgraphics: {dgraphics}</pre>\n\n"

    return AnimeSched(day, curr_page, has_next, res)


def __gen_kbd(page: int, has_next: bool, udata: str) -> InlineKeyboardMarkup:
    builder = InlineKeyboardBuilder()
    if page > 1:
        builder.button(
            text="Prev",
            callback_data=f"{CALLBACK_ID}:{page - 1}:{udata}"
        )

    if has_next:
        builder.button(
            text="Next",
            callback_data=f"{CALLBACK_ID}:{page + 1}:{udata}"
        )

    return builder.as_markup()


#
# Messages
#
@r.message(filters.Command("anime_schedule"))
async def cmd_anime_sched_fn(msg: Message) -> None:
    try:
        _, arg = utils.tok_1(msg.text, " ")

        arg = arg.lower()
        if len(arg) > 0 and arg not in utils.days:
            await msg.reply(err(f"invalid day name: {arg}"))
            return

        res = await __get_schedule(day=arg)
        await msg.reply(
            text=res.result,
            reply_markup=__gen_kbd(res.page, res.has_next, arg)
        )
    except Exception as e:
        await msg.reply(err(e))


async def run_cb(bot: Bot, msg: Message, udata: str) -> None:
    page, day = utils.tok_1(udata, ":")

    res = await __get_schedule(day, page)
    await bot.edit_message_text(
        chat_id=msg.chat.id, message_id=msg.message_id,
        text=res.result
    )

    await bot.edit_message_reply_markup(
        chat_id=msg.chat.id,
        message_id=msg.message_id,
        reply_markup=__gen_kbd(res.page, res.has_next, res.day)
    )
