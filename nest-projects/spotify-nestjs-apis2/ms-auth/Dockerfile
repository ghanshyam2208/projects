# Use the official Node.js image from the Docker Hub
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

COPY ./ms-auth/pnpm-lock.yaml .

COPY ./ms-auth/package.json .

RUN npm install -g pnpm

RUN pnpm install

COPY ./ms-auth .

CMD ["pnpm", "start"]
# CMD [ "sh" ]