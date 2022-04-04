ARG IMAGE="library/golang"
ARG VERSION="latest"

FROM "${IMAGE}:${VERSION}"

MAINTAINER macwinnie <dev@macwinnie.me>

# environmental variables
ENV TERM            "xterm"
ENV DEBIAN_FRONTEND "noninteractive"
ENV TIMEZONE        "Europe/Berlin"
ENV SET_LOCALE      "de_DE.UTF-8"
ENV WORKDIR         "/project"

# copy all relevant files
COPY src/ "${WORKDIR}"
COPY files/ "/install.d"

RUN chmod a+x /install.d/install.sh && \
    /install.d/install.sh && \
    rm -rf /install.d

WORKDIR "${WORKDIR}"

# run on every (re)start of container
# ENTRYPOINT [ "entrypoint" ]
CMD [ "sleep infinity" ]
