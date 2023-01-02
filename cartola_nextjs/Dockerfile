FROM node:16.16.0-slim

COPY . /home/node/app

WORKDIR /home/node/app

RUN cp .env.example .env.local  
RUN npm install

CMD [ "npm", "run", "dev" ]