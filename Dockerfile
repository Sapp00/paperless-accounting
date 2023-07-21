FROM python:3

LABEL maintainer="support@heymann.dev"
LABEL version="0.1"
LABEL description="A simple accounting solution based on paperless"

RUN pip install Django

RUN mkdir /accounting/
COPY src /accounting/src

EXPOSE 8000

CMD [ "python", "./mange.py runserver"]