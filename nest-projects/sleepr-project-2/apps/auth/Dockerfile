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

# Install dependencies using pnpm
RUN pnpm install

# Copy the rest of the application code into the container
COPY . .

# Build the application using pnpm
RUN pnpm run build

# # Use the Node.js Alpine image as the base for production
# FROM node:alpine AS production

# # Set the NODE_ENV environment variable to production
# ARG NODE_ENV=production
# ENV NODE_ENV=${NODE_ENV}

# # Set the working directory inside the container
# WORKDIR /usr/src/app

# # Copy package.json from the current directory into the container
# COPY package.json ./

# RUN sed -i 's/0.0.0.0:9229/0.0.0.0:9231/g' /usr/src/app/package.json

# # Copy pnpm-lock.yaml from the current directory into the container
# COPY pnpm-lock.yaml ./

# # Install pnpm globally using npm
# RUN npm install -g pnpm

# # Install production dependencies using pnpm
# RUN pnpm install --prod

# # Copy the built application from the development stage
# COPY --from=development /usr/src/app/dist ./dist

RUN sed -i 's/0.0.0.0:9229/0.0.0.0:9232/g' /usr/src/app/package.json

# Command to start the application
CMD ["node", "dist/apps/auth/main"]
