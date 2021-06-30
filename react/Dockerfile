FROM node:14.17.1-alpine as build
WORKDIR /react-app
ENV PATH /react-app/node_modules/.bin:$PATH
COPY ./package.json /react-app/
RUN npm i --silent
COPY . /react-app
CMD ["npm","start"]