from sanic import Sanic
from sanic_cors import CORS

from src.v1.setup_route import setup_route

app = Sanic(__name__)
CORS(app)
setup_route(app)

if __name__ == '__main__':
    try:
        app.run(port=8000)
    except KeyboardInterrupt:
        print('\n\nInterrupted by user.\n')
        exit(0)