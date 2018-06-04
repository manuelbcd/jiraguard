FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./main /go/src/github.com/manuelbcd/jiraguard/main
WORKDIR /go/src/github.com/manuelbcd/jiraguard/main

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	main; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080