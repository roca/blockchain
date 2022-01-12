# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM node:latest

RUN npm install -g choco
RUN npm install -g truffle

CMD ["tail", "-f", "/dev/null"]