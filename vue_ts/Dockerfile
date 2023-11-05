FROM node:16.16

WORKDIR /vue_ts
COPY . /vue_ts

RUN npm install
    
EXPOSE 8080

CMD sh -c "npm run serve"