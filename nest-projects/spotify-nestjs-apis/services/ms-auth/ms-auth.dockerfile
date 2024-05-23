# Use the official Node.js image from the Docker Hub
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

COPY pnpm-lock.yaml .

COPY package.json .
COPY pnpm-workspace.yaml .

RUN mkdir services
RUN mkdir services/ms-auth

COPY ./services/ms-auth/package.json ./services/ms-auth

RUN npm install -g pnpm

RUN pnpm --filter "ms-auth" install 


COPY ./services/ms-auth ./services/ms-auth

# Start a bash shell by default for debugging
# CMD ["sh"]
# Start the application
CMD ["pnpm", "--filter", "ms-auth", "start"]

