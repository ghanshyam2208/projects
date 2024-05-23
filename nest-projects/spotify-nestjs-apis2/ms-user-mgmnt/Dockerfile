# Use the official Node.js image from the Docker Hub
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

COPY ./ms-user-mgmnt/pnpm-lock.yaml .

COPY ./ms-user-mgmnt/package.json .

RUN npm install -g pnpm

RUN  npm install -g @nestjs/cli

RUN pnpm install

COPY ./ms-user-mgmnt .

CMD ["pnpm", "start"]
# CMD [ "sh" ]
