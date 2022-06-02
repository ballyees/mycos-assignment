from sanic import Sanic
from sanic.blueprints import Blueprint
from .blueprints.employee import bp as bp_employee
def setup_route(app: Sanic):
    
    bpg = Blueprint.group([bp_employee], url_prefix="/", version="v1")
    app.blueprint(bpg)
    