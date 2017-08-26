# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:v0.9.3 as builder

RUN mkdir -p $GOPATH/src/github.com/klauern/shellchecker
WORKDIR $GOPATH/src/github.com/klauern/shellchecker

ADD . .
RUN dep ensure
RUN buffalo build --static -o /bin/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache curl
RUN curl -O https://storage.googleapis.com/shellcheck/shellcheck-latest.linux.x86_64.tar.xz
RUN tar xvf shellcheck-latest.linux.x86_64.tar.xz
RUN mv shellcheck-latest/shellcheck /bin/
RUN rm -r shellcheck*


# Comment out to run the binary in "production" mode:
# ENV GO_ENV=production

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 3000

# Comment out to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD /bin/app
