FROM node:latest
RUN apk add gcompat
RUN wget -qO- https://get.pnpm.io/install.sh | ENV="$HOME/.shrc" SHELL="$(which sh)" sh -
RUN source /root/.shrc && pnpm env use --global 20
