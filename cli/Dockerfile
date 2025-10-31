FROM ubuntu:22.04

WORKDIR /app

RUN apt-get update && apt-get install curl gnupg jq -y

RUN mkdir -p /etc/apt/keyrings && \
    curl -fsSL https://repo.charm.sh/apt/gpg.key | gpg --dearmor -o /etc/apt/keyrings/charm.gpg && \
    echo "deb [signed-by=/etc/apt/keyrings/charm.gpg] https://repo.charm.sh/apt/ * *" | tee /etc/apt/sources.list.d/charm.list && \
    apt update && apt install gum

COPY ./cli.sh /app/cli.sh

CMD /app/cli.sh
