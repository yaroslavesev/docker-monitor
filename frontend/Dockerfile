FROM node:19-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install -g @angular/cli && npm install

RUN npm install bootstrap @ng-bootstrap/ng-bootstrap

COPY . .

EXPOSE 4200

CMD ["npm", "run", "start"]
