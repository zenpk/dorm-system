FROM node:16-alpine

WORKDIR /app

COPY . .

RUN npm config set registry https://registry.npmmirror.com
RUN npm install
RUN npm run build

ENV NODE_ENV production
ENV NEXT_TELEMETRY_DISABLED 1

EXPOSE 3000

CMD ["npm", "run", "start"]
