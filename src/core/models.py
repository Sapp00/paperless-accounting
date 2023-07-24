from django.db import models

import json

class Singleton(models.Model):

    class Meta:
        abstract = True

    def save(self, *args, **kwargs):
        self.pk = 1
        super(Singleton, self).save(*args, **kwargs)

    def delete(self, *args, **kwargs):
        pass

    @classmethod
    def load(cls):
        obj, _ = cls.objects.get_or_create(pk=1)
        return obj

class DynamicSettings(Singleton):
    expense_tag = models.CharField(max_length=32)
    income_tag = models.CharField(max_length=32)


class PaperlessDocument(object):
    def __init__(self, id, correspondent, title, content, tags, created_date):
        self.id = id
        self.correspondent = correspondent
        self.title = title
        self.content = content
        self.tags = tags
        self.created_date = created_date
    
    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, 
                    sort_keys=True, indent=4)