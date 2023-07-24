from django.db import models

# Create your models here.
class Expense(models.Model):
    id = models.IntegerField(unique=True, primary_key=True)
    price = models.DecimalField(decimal_places=2, max_digits=12)
    paid = models.BooleanField()
    paidDate = models.DateField("date published")
