FROM node:24.22.1-bookworm-slim@sha256:be8204d4bd327363357f532187b752ff62759c31bedf98e9e18fc08c9cbd605c

WORKDIR /app

COPY . .

RUN npm install && npm run build

ENTRYPOINT ["npm", "run", "server"]
