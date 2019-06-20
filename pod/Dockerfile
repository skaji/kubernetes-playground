ARG github_user=skaji
ARG github_repo=kubernetes-playground/pod
ARG version=dummy
ARG commitish=dummy
ARG port=8080
ARG timezone=Asia/Tokyo

FROM golang as builder
ARG github_user
ARG github_repo
ARG version
ARG commitish
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/$github_user/$github_repo
COPY ./ ./
RUN test -f go.mod && export GO111MODULE=on; go build -o app -ldflags "-X main.version=$version -X main.commitish=$commitish"

FROM alpine
ARG github_user
ARG github_repo
ARG timezone
ARG port
RUN set -exo pipefail; \
  apk add --update --no-cache tzdata ca-certificates; \
  cp /usr/share/zoneinfo/$timezone /etc/localtime; \
  echo $timezone > /etc/timezone; \
  apk del tzdata; \
  true
COPY --from=builder /go/src/github.com/$github_user/$github_repo/app /app
EXPOSE $port
ENTRYPOINT ["/app"]
