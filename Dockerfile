FROM alpine:latest
<<<<<<< HEAD
RUN apk --update add ca-certificates \
                     mailcap \
                     curl \
                     jq

COPY healthcheck.sh /healthcheck.sh
RUN chmod +x /healthcheck.sh  # Make the script executable
=======
>>>>>>> a69839c7 (docker uses alpine)

HEALTHCHECK --start-period=2s --interval=5s --timeout=3s \
    CMD /healthcheck.sh || exit 1

VOLUME /srv
EXPOSE 80

COPY docker_config.json /.filebrowser.json
COPY filebrowser /filebrowser

ENTRYPOINT [ "/filebrowser" ]