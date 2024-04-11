FROM node:20.10.0-alpine

WORKDIR /front

COPY package.json yarn.lock ./

RUN yarn install

COPY . .


CMD ["yarn", "dev"]