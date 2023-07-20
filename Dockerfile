FROM alpine:latest

HEALTHCHECK --start-period=2s --interval=5s --timeout=3s \
  CMD curl -f http://localhost/health || exit 1

VOLUME /srv
EXPOSE 80

COPY docker_config.json /.filebrowser.json
COPY filebrowser /filebrowser

ENTRYPOINT [ "/filebrowser" ]