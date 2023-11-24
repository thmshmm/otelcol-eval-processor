FROM node:18-alpine

RUN npm install -g @mockoon/cli@5.1.0
COPY ./mock.json ./mock.json

# Install curl for healthcheck and tzdata for timezone support.
RUN apk --no-cache add curl tzdata

# Do not run as root.
RUN adduser --shell /bin/sh --disabled-password --gecos "" mockoon
RUN chown -R mockoon ./mock.json
USER mockoon

EXPOSE 8086

ENTRYPOINT ["mockoon-cli","start","--disable-log-to-file","--data","./mock.json","--port","8086"]

# Usage: docker run -p <host_port>:<container_port> mockoon-test
