FROM node:20.6

RUN mkdir /app
COPY package.json /app
COPY tsconfig.json /app
COPY frontend-nodejs.ts /app

WORKDIR /app
RUN npm install
CMD ["npm", "start"]
