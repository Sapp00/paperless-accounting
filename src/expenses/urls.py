from django.urls import path

from . import views

urlpatterns = [
    path("", views.index, name="index"),
    path("all_expenses", views.allExpenses, name="allExpenses"),
    path("<int:expense_id>/", views.detail, name="detail"),

]