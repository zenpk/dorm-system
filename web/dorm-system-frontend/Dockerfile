FROM node:16-alpine

WORKDIR /app

COPY . .

RUN npm ci

RUN npm run build

ENV NODE_ENV production
ENV NEXT_TELEMETRY_DISABLED 1

EXPOSE 3000

CMD ["npm", "run", "start"]
