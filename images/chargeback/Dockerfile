FROM alpine

RUN apk update
RUN apk add ca-certificates bash
# add pod data collector binary
ADD ./bin/chargeback /usr/local/bin/chargeback

ENTRYPOINT ["chargeback"]
CMD ["start"]

USER 3001
