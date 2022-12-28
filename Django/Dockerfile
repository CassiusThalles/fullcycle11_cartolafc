FROM python:3.10.2-slim

RUN useradd -ms /bin/bash python
# RUN cp .env.example .env 
# RUN pip install pipenv 
# RUN pipenv install 
# RUN pipenv run python manage.py migrate 
# RUN pipenv run python manage.py loaddata initial_data 

USER python

WORKDIR /home/python/app

ENV PIPENV_VENV_IN_PROJECT=True

CMD [ "./.docker/start.dev.sh" ]