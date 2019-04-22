FROM golang
RUN mkdir app/ && cd app/
WORKDIR app/
COPY . .
RUN go build main.go
CMD ["./main"]
EXPOSE 8085
