FROM node:23-slim

WORKDIR /app

COPY . .

RUN npm install && npm run build

ENTRYPOINT ["npm", "run", "server"]
