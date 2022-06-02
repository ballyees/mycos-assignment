from sanic.blueprints import Blueprint
from sanic.response import json
from ..database.employee import Qeury
from datetime import datetime as dt

bp = Blueprint('employee', url_prefix='/pvd')
time_format = '%Y-%m-%d'

@bp.route('/', methods=["GET"])
async def get_all_employee(request):
    data = Qeury.get_all_employee()
    reponse = [{} for i in range(len(data))]
    now = dt.now()
    for i, d in enumerate(data):
        reponse[i]['Employee'] = d
        reponse[i]['TotalPVD'] = calculate_total_pvd(d, now)
        reponse[i]['TotalMonthPVD'] = calculate_total_month_pvd(d, now)
    return json(reponse)

def calculate_pvd(employee, month, paid_rate):
    return ((employee['Salary'] * paid_rate / 100) * month) + ((employee['Salary'] * employee['PvdFundRate'] / 100) * month)

def calculate_total_pvd(employee, time_now):
    total_pvd = 0
    t =  dt.strptime(employee['StartDate'], time_format)
    diff = time_now - t
    year = diff.days // 365
    month = diff.days % 365 // 30
    if year >= 5: # over 5 year
        total_pvd += calculate_pvd(employee, 12*(year-5)+month, 80)
        year = 5
        month = 0
    if year >= 3: # less than 5 year
        total_pvd += calculate_pvd(employee, 12*(year-3)+month, 50)
        year = 3
        month = 0
    if year >= 1: # less than 3 year
        total_pvd += calculate_pvd(employee, 12*(year-1)+month, 30)
        year = 1
        month = 0
    total_pvd += calculate_pvd(employee, max((12*year)+month-3, 0), 10) # less than 1 year
    return total_pvd

def calculate_total_month_pvd(employee, time_now):
    total_pvd = 0
    t =  dt.strptime(employee['StartDate'], time_format)
    diff = time_now - t
    year = diff.days // 365
    month = diff.days % 365 // 30
    total_month = (12*year + month) - 3 
    if total_month <= 0:
        return 0
    if year >= 5: # over 5 year
        return calculate_pvd(employee, total_month, 80)
    elif year > 3: # less than 5 year
        return calculate_pvd(employee, total_month, 50)
    elif year > 1: # less than 3 year
        return calculate_pvd(employee, total_month, 30)
    else: # less than 1 year
        return calculate_pvd(employee, total_month, 10)