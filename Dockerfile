FROM scratch

ARG TELEGRAM_API_KEY

LABEL maintainer="oshalygin@gmail.com"
LABEL description="Telegram StravaBot"

ADD tls-ca-bundle.pem /etc/ssl/certs/

EXPOSE 8080
ADD main /

CMD ["/main"]
