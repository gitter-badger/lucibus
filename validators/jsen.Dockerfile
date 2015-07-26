FROM node:12.7
RUN npm install jsen@0.6.0

ADD validators/jsen.js /code/jsen.js
ADD sample.json /code/sample.json
ADD schema.json /code/schema.json

CMD node /code/jsen.js
