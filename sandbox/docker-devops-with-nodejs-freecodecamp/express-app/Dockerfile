FROM node:18-alpine
WORKDIR /usr/src/app
COPY package.json .
COPY package-lock.json .

RUN if [ "$NODE_ENV" = "production" ]; \
        then npm install --only=production; \
        else npm install; \
    fi

COPY . .
CMD ["node", "index.js"]