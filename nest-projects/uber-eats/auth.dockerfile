# Use the Node.js Alpine image as the base for development
FROM node:alpine AS development

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json from the current directory into the container
COPY package.json ./

# Copy pnpm-lock.yaml from the current directory into the container
COPY pnpm-lock.yaml ./

# Install pnpm globally using npm
RUN npm install -g pnpm

# # Install dependencies using pnpm
# RUN pnpm --filter "auth-ms" install

# Copy the rest of the application code into the container
COPY . .

# Build the application using pnpm
# RUN pnpm run build
