FROM node:20.6

RUN mkdir /app
COPY package.json /app
COPY frontend-nodejs.js /app

WORKDIR /app
RUN yarn install
CMD ["yarn", "start"]
