# base image
FROM node:14.7.0-alpine as builder

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY ./backend/package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY ./backend/ ./

# build app for production with minification
RUN npm run build

# ############# #
# NGINX builder #
# ############# #
FROM nginx:1.21-alpine

# remove default nginx index
RUN rm -rf /usr/share/nginx/html/*
WORKDIR /usr/share/nginx/html

RUN mkdir front
COPY ./frontend/ ./front/

RUN mkdir back
COPY --from=builder /app/dist ./backend/

# first remove the current existing config
RUN rm /etc/nginx/conf.d/default.conf
# then copy the config file to the container
COPY nginx.conf /etc/nginx/conf.d
