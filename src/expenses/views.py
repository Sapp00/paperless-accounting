from django.shortcuts import render
from django.http import HttpResponse, JsonResponse, HttpResponseBadRequest
from django.conf import settings
from django.template import loader

from core.models import DynamicSettings, PaperlessDocument

import requests
import urllib.parse
from urllib3.exceptions import InsecureRequestWarning
import json

import warnings

# Create your views here.

def allExpenses(request):
    try:
        # load tag
        dynamicSettings = DynamicSettings.load()
        expense_tag = dynamicSettings.expense_tag

        # send request
        requests.packages.urllib3.disable_warnings(InsecureRequestWarning)
        response = requests.get(
            urllib.parse.urljoin(settings.PAPERLESS_URL, "/api/documents/?query=tag:"+expense_tag),
            headers = {'Authorization': 'Token '+settings.PAPERLESS_AUTH_TOKEN },
            verify = not settings.PAPERLESS_UNSAFE_SSL
        )

        if settings.DEBUG:
            response.raise_for_status()

        all_expenses = [ PaperlessDocument(k['id'], k['correspondent'], k['title'], k['content'], k['tags'], k['created_date']) for k in response.json()['results']]

        all_expenses_json = { 
                "results": [ e.toJSON() for e in all_expenses ]
            }

        return JsonResponse(all_expenses_json)

    except requests.HTTPError as http_err:
        print(f'HTTP error occured: {http_err}')
    except Exception as err:
        print(f'Unexpected error occurred: {err}')

    return HttpResponseBadRequest('An error occured')

def index(request):
    try:
        # load tag
        dynamicSettings = DynamicSettings.load()
        expense_tag = dynamicSettings.expense_tag

        # send request
        requests.packages.urllib3.disable_warnings(InsecureRequestWarning)
        response = requests.get(
            urllib.parse.urljoin(settings.PAPERLESS_URL, "/api/documents/?query=tag:"+expense_tag),
            headers = {'Authorization': 'Token '+settings.PAPERLESS_AUTH_TOKEN },
            verify = not settings.PAPERLESS_UNSAFE_SSL
        )

        if settings.DEBUG:
            response.raise_for_status()

        all_expenses = [ PaperlessDocument(k['id'], k['correspondent'], k['title'], k['content'], k['tags'], k['created_date']) for k in response.json()['results']]

        expense_chart_dict = dict()
        expense_chart_paid_dict = dict()
        expense_sum = 0
        paid_sum = 0
        for e in all_expenses:
            # TODO: change created date! needs to be based on paid_date which is retrieved from the database
            e_price = ord(e.title[0])*20
            e_paid = e_price if e_price % 3 != 0 else 0
            e_paid_date = e.created_date
            e_date = e.created_date

            paid_sum += e_paid
            expense_sum += e_price
            if e_date in expense_chart_dict:
                expense_chart_dict[ e_date ]['value'] = expense_sum
            else:
                expense_chart_dict[ e_date ] = {'date': e_date, 'category': 'expense', 'value': expense_sum}

            if e_paid_date in expense_chart_paid_dict:
                expense_chart_paid_dict[ e_paid_date ]['value'] = paid_sum
            else:
                expense_chart_paid_dict[ e_paid_date ] = {'date': e_paid_date, 'category': 'payment', 'value': paid_sum}

        # merge expenses+income
        expense_chart_data = list(expense_chart_dict.values()) + list(expense_chart_paid_dict.values())

        # render result
        template = loader.get_template("expenses/index.html")
        context = {
            "all_expenses": all_expenses,
            "expense_chart_data": expense_chart_data
        }
        return HttpResponse(template.render(context, request))
        
    except requests.HTTPError as http_err:
        print(f'HTTP error occured: {http_err}')
    except Exception as err:
        print(f'Unexpected error occurred: {err}')

def detail(request, expense_id):
    return HttpResponse("You're looking at expense %s." % expense_id)