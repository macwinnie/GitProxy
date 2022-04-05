# Git Proxy

This repository is meant to provide a Git proxy for environments, where Git via HTTP/S and / or SSH is restricted.

You'll need an user directory like a LDAP server, from where your users will be authorized to use the WebApp and also share Git projects with LDAP groups.

## Development

For development, the `docker-compose.yml` provides a Go-environment to work with local `src` files.

For production, you should use `macwinnie/gitproxy:latest` image.

## Environmental variables

| env                   | default               | change recommended | description |
| --------------------- | --------------------- |:------------------:| ----------- | `WEB_PORT`            | `80`                  | no                 | On which port should the webserver within the container be listening? For port-binding, it is recommended to change the port not within the container. |

## Licence

This project is published unter [CC BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/) license.

## last built

0000-00-00 00:00:00
