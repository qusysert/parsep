FROM node:12 as builder
WORKDIR /app

RUN mkdir -p app/node_modules && chown -R node:node /app
COPY . .
USER node
RUN npm install && npm run
CMD [ "node", "app.js" ]

