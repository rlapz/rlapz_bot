import aiohttp
from aiohttp import ClientResponse
from datetime import datetime
import pytz

tz = pytz.timezone("Asia/Jakarta")
days = (
    'monday',
    'tuesday',
    'wednesday',
    'thursday',
    'friday',
    'saturday',
    'sunday',

    ##### anime #####
    'unknown',
    'other',
)


def get_current_day() -> str:
    idx = datetime.now(tz).weekday()
    return days[idx]


def err(arg: str) -> str:
    return f"<pre>Error occured: {arg}</pre>"


def tok_1(arg: str, sep: str) -> (str, str):
    ret: str = ""
    pos: int = 0
    for i in arg:
        if i == sep:
            pos += 1
            break
        ret += i
        pos += 1

    return (ret.strip(" "), arg[pos:].strip(" "))


async def http_request_json(url: str) -> (ClientResponse, object()):
    async with aiohttp.ClientSession() as sess:
        async with sess.get(url) as resp:
            return (resp, await resp.json())


def join_list_dict(arg, key: str) -> str:
    _list = []
    for x in arg:
        _list.append(x[key])

    return ', '.join(_list)
