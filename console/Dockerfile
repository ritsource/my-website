FROM node:10.15.3-alpine as builder

WORKDIR /usr/src

COPY ./package.json .
RUN yarn install

COPY . .

RUN yarn run build


# Serve Using Nginx (Production)
FROM nginx

EXPOSE 3000

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html