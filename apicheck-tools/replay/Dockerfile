FROM node:14-alpine

WORKDIR /app
COPY . /app

RUN npm install && npm run build && npm install -g
RUN npm prune --production

ENTRYPOINT ac-replay
