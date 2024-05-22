# Use the official Node.js image from the Docker Hub
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

COPY pnpm-lock.yaml .

COPY package.json .
COPY pnpm-workspace.yaml .

RUN mkdir services
RUN mkdir services/ms-users-mgmt

COPY ./services/ms-users-mgmt/package.json ./services/ms-users-mgmt

RUN npm install -g pnpm

RUN pnpm --filter "ms-users-mgmt" install 


COPY ./services/ms-users-mgmt ./services/ms-users-mgmt

# Start a bash shell by default for debugging
# CMD ["sh"]
# Start the application
CMD ["pnpm", "--filter", "ms-users-mgmt", "start"]

