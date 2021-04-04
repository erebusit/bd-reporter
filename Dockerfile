FROM circleci/golang:1.14
RUN go get -u -v github.com/erebusit/bd-reporter
RUN bd-reporter --help
ENTRYPOINT ["bd-reporter"]
