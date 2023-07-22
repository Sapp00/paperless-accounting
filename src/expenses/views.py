from django.shortcuts import render
from django.http import HttpResponse

from django.conf import settings
import requests

# Create your views here.

def index(request):
    
    return HttpResponse("Hello world. You're at the pools")
    