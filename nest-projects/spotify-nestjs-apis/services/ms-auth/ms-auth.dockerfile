# Use the official Node.js image from the Docker Hub
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

COPY pnpm-lock.yaml .

COPY ./services/ms-auth/package.json .

RUN npm install -g pnpm

RUN pnpm --filter "ms-auth" install 


COPY ./services/ms-auth .

# Start a bash shell by default for debugging
# CMD ["sh"]
# Start the application
CMD ["pnpm", "start"]

