FROM nginx:1.15.3-alpine

RUN rm /etc/nginx/conf.d/*
COPY etc/nginx/nginx.conf /etc/nginx/
COPY etc/nginx/conf.d/ /etc/nginx/conf.d/

RUN rm /usr/share/nginx/html/*
