# build stage
FROM node:12-stretch AS build

WORKDIR /build

ARG src=fake_server

COPY --chown=node:node ${src}/package-lock.json ${src}/package.json ./

RUN npm ci --production

COPY ${src} .

# runtime stage
FROM alpine:3.10

RUN apk add --update nodejs-current

RUN addgroup -S node && adduser -S node -G node

USER node

RUN mkdir /home/node/code

WORKDIR /home/node/code

COPY --from=build --chown=node:node /build .

CMD ["node", "index.js"]