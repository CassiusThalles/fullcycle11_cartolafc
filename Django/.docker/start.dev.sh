#!/bin/bash

if [ ! -f ".env" ]; then
  cp .env.example .env
fi

python -m pip install pipenv

python -m pipenv install

python -m pipenv run python manage.py migrate

python -m pipenv run python manage.py loaddata initial_data

python -m pipenv run python manage.py runserver 0.0.0.0:8000

# tail -f /dev/null