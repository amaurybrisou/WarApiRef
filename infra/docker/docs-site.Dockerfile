FROM nginx:1.27-alpine

COPY docs/site /usr/share/nginx/html
