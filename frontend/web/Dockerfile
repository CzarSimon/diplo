FROM mhart/alpine-node:10 as build
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build

FROM nginx:1.13.8-alpine as run
WORKDIR /etc/nginx
COPY --from=build /app/build html/
COPY nginx.conf nginx.conf
RUN nginx -t
