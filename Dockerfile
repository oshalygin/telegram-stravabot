FROM alpine:3.7

ARG TELEGRAM_API_KEY

LABEL maintainer="oshalygin@gmail.com"
LABEL description="Telegram StravaBot"

EXPOSE 8080
ADD main /

CMD ["/main"]
