FROM golang:1.19-alpine

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -o crypto-currency.api .

EXPOSE 8085

CMD [ "./crypto-currency.api" ]