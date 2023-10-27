FROM node:16.16

WORKDIR /vue
COPY . /vue

RUN apt-get update -qq \
    && yarn global add @vue/cli
    
EXPOSE 8080

CMD sh -c "yarn && yarn run serve"