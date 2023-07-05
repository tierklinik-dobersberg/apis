# Build the frontend
FROM node:16

WORKDIR /src

COPY ./ ./
RUN npm install
