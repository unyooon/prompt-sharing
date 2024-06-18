FROM node:18-slim

WORKDIR /app

RUN npm install -g pnpm

COPY . /app

RUN rm -rf .git/
RUN rm -rf node_modules

RUN pnpm install
RUN pnpm build

# RUN npx nuxi upgrade --force

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]