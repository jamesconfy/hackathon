FROM golang:1.19-alpine3.16

RUN mkdir /cheque_deposit

COPY . /cheque_deposit

WORKDIR /cheque_deposit

LABEL Name=cheque_deposit Version=0.0.1

RUN go build -o cheque_deposit

EXPOSE  1010

CMD [ "./cheque_deposit", "--m=true" ]