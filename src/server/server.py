import logging
import multiprocessing as mp
import struct
import threading
from argparse import ArgumentParser
from datetime import datetime
from time import sleep

import tornado.ioloop
import tornado.web
import tornado.websocket
from aeronpy import Context
from perspective import Table, PerspectiveTornadoHandler, PerspectiveManager

logging.basicConfig(level=logging.INFO, format="%(asctime)s - %(process)d %(levelname)s: %(message)s")


class GeminiOrderBookDataSource(object):

    def __init__(self, symbol, data_queue):
        """A datasource that interfaces with the Gemini Websockets API to
        receive live order book data and submits it to a queue in order
        to update the Perspective table."""
        self.symbol = symbol
        self.data_queue = data_queue
        self.url = "wss://api.sandbox.gemini.com/v1/marketdata/{}?bids=true&offers=true&trades=false".format(
            self.symbol)

        parser = ArgumentParser()
        parser.add_argument('-p', '--prefix', default="aeron-dir")
        parser.add_argument('-c', '--channel', default='aeron:udp?endpoint=localhost:40123')
        parser.add_argument('-s', '--stream_id', type=int, default=1)

        args = parser.parse_args()
        self.aeron_dir = args.prefix
        self.channel = args.channel
        self.stream_id = args.stream_id

    def start(self):
        """Make the API connection."""
        # logging.info("Connecting to Gemini for {} order book".format(self.symbol))
        # self.ws = websocket.WebSocketApp(self.url, on_message=self.on_message)
        # self.ws.run_forever(sslopt={"cert_reqs": ssl.CERT_NONE})

        context = Context(
            aeron_dir=self.aeron_dir,
            new_subscription_handler=lambda *args: print(f'new subscription {args}'),
            available_image_handler=lambda *args: print(f'available image {args}'),
            unavailable_image_handler=lambda *args: print(f'unavailable image {args}'))

        subscription = context.add_subscription(self.channel, self.stream_id)
        while True:
            fragments_read = subscription.poll(self.on_message)
            if fragments_read == 0:
                eos_count = subscription.poll_eos(lambda *args: print(f'end of stream: {args}'))
                if eos_count > 0:
                    break

            sleep(0.1)

    # def format_msg(self, msg):
    #     """Given a message from the Gemini order book, format it properly
    #     for the Perspective table."""
    #     formatted = []
    #     timestamp = msg.get("timestamp")
    #
    #     if timestamp:
    #         timestamp = datetime.fromtimestamp(timestamp)
    #     else:
    #         timestamp = datetime.now()
    #
    #     for event in msg["events"]:
    #         event["symbol"] = self.symbol
    #         event["timestamp"] = timestamp
    #         event["price"] = float(event["price"])
    #         event["remaining"] = float(event["remaining"])
    #         event["delta"] = float(event["delta"])
    #
    #         formatted.append(event)
    #
    #     return formatted

    def on_message(self, msg):
        """Format the message from Gemini and add it to the queue so
        the data updater thread can pick it up and apply it to the table."""
        if msg is None:
            logging.warn("Gemini API connection closed for symbol {}".format(self.symbol))
            return

        memory = memoryview(msg)
        # b = bytes(memory)

        offset = 0

        event = {}

        def read_string(off):
            str_len = struct.unpack_from("i", memory, off)[0]
            off += 4

            val = str(memory[off: off + str_len], 'ascii')
            # print(val)
            off += str_len
            return off, val

        offset, symbol = read_string(offset)
        event["symbol"] = symbol

        offset, timestamp = read_string(offset)
        # event["timestamp"] = datetime.fromtimestamp(timestamp)
        event["timestamp"] = datetime.now()

        offset, side = read_string(offset)
        event["side"] = side

        offset, action = read_string(offset)
        event["type"] = action

        offset, order_id = read_string(offset)
        event["order_id"] = order_id

        offset, client_id = read_string(offset)
        event["client_id"] = client_id

        # print(symbol, side, action, order_id, client_id)

        price, initial, size, filled, post_only, ioc, delta, signed_delta, diff_price, diff_size, tick_size, contract_size = struct.unpack_from(
            '4l2i2l2i2d', memory[offset:])

        event["price"] = float(price)

        event["reason"] = client_id

        event["delta"] = float(delta)

        event["remaining"] = float(size)
        #
        # print(price, initial, size, filled, post_only, ioc, delta, signed_delta, diff_price, diff_size, tick_size,
        #       contract_size)

        # msg = json.loads(msg)
        # self.data_queue.put(self.format_msg(msg))
        self.data_queue.put([event])


MANAGER = PerspectiveManager()
PSP_LOOP = tornado.ioloop.IOLoop()

ORDER_BOOK = Table({
    "symbol": str,
    "type": str,
    "reason": str,
    "side": str,
    "price": float,
    "delta": float,
    "remaining": float,
    "timestamp": datetime
}, limit=5000)


def perspective_thread():
    """Run Perspective on a separate thread using a Tornado IOLoop,
    which improves concurrent performance with multiple clients."""
    MANAGER.set_loop_callback(PSP_LOOP.add_callback)
    MANAGER.host_table("order_book", ORDER_BOOK)
    PSP_LOOP.start()


def fetch_data(table, data_queue):
    """Wait for the datasource to add new data to the queue, and call
    table.update() using the IOLoop's add_callback method in order to call
    the operation on the Perspective thread."""
    while True:
        data = data_queue.get()
        print(data)
        PSP_LOOP.add_callback(table.update, data)


def start():
    """Set up the server - we use a queue to manage the flow of data from the
    datasource to the Table. There are two processes: the main process which
    runs the Tornado server and two sub-threads, one thread for Perspective
    and another thread to fetch data from the queue, and the subprocess which
    runs the datasource and submits data to the queue in order to transfer it
    between processes."""
    orders_queue = mp.Queue()

    # The thread that fetches data from the queue and calls table.update
    order_fetcher_thread = threading.Thread(target=fetch_data, args=(ORDER_BOOK, orders_queue))
    order_fetcher_thread.daemon = True
    order_fetcher_thread.start()

    # The thread that runs Perspective
    psp_thread = threading.Thread(target=perspective_thread)
    psp_thread.daemon = True
    psp_thread.start()

    # The process that runs the datasource
    orders = GeminiOrderBookDataSource("ethusd", orders_queue)
    orders_process = mp.Process(target=orders.start)
    orders_process.start()

    app = tornado.web.Application([
        (
            r"/websocket",
            PerspectiveTornadoHandler,
            {"manager": MANAGER, "check_origin": True},
        ),
    ])

    # Tornado listens on the main process
    app.listen(8081)
    app_loop = tornado.ioloop.IOLoop()
    app_loop.make_current()
    app_loop.start()


if __name__ == "__main__":
    start()
