from aiohttp import web
import asyncio
import orjson

async def handle_ping(request):
    return web.Response(text="pong")

async def handle_bidder(request):
    if request.content_type != "application/json" or not request.can_read_body:
        return web.Response(status=400)
    body = await request.text()
    br = orjson.loads(body)
    return web.Response(status=204)

def main():
    app = web.Application()
    app.add_routes([
        web.get('/ping', handle_ping),
        web.post('/bidder', handle_bidder),
        ])
    loop = asyncio.new_event_loop()
    #loop.set_debug(True)
    print("NOBIDDER PYTHON-AIOHTTP Running on http://127.0.0.1:8080/")
    web.run_app(app, host="127.0.0.1", port=8080, loop=loop)

if __name__ == '__main__':
    main()
