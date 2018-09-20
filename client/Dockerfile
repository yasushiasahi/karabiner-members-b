FROM node:10.9.0-alpine

RUN apk add curl

COPY . /client

WORKDIR /client

RUN ls -al
RUN npm install

CMD ["npm","run","serve"]

EXPOSE 8080
