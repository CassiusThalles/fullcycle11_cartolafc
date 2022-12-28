#!/bin/bash

cp .env.example .env

python -m pip install pipenv

python -m pipenv install

python -m pipenv run python manage.py migrate

python -m pipenv run python manage.py loaddata initial_data

tail -f /dev/null