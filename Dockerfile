FROM node:latest
WORKDIR /build
COPY . /build
RUN yarn

EXPOSE 3000
CMD ["node", "index.js"]
