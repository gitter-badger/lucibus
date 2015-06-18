# caido


## Development

## Design
Initial structure based off of [`react-starter`](https://github.com/webpack/react-starter/tree/48cecfcd3a528ceefdd3d68b4e0f05fffbedac8e)

We are using [webpack](https://github.com/webpack/webpack) to bundle up everything
and compile. So if you wanna add a stylesheet don't slap some `<link href="">...`
crap in a HTML file somewhere. Instead just import it in your componenent
which uses it, after installing with `npm`.

```
import "bootstrap/dist/css/bootstrap.css";
```

Also all custom styles should be completely local. Check out
[this demo/instruction site](https://css-modules.github.io/webpack-demo/)
for examples.

### Requirements

1. Install [Docker](https://get.docker.com/)
2. Install [Docker Compose](https://docs.docker.com/compose/)

#### Rebuilding

When you change a dependency in `package.json` you need to rebuild

```
docker-compose -f docker-compose/common.yml build web

# sometimes you need to remove all the old docker containers after, if it doesn't
# pick up the new changes
docker rm -f $(docker ps -a -q)
```

#### For a mac install
Use the [docker-osx-dev](https://github.com/brikis98/docker-osx-dev)
project to get proper watch events. Install everything by it's commands

Then run `docker-osx-dev` from the root directory of this project

### Development server

```
docker-compose -f docker-compose/dev-server.yml up
open http://dockerhost:8080/webpack-dev-server/
```

The configuration is `webpack-dev-server.config.js`.

It automatically recompiles when files are changed. When a hot-replacement-enabled file is changed (i. e. stylesheets or React components) the module is hot-replaced. If Hot Replacement is not possible the page is refreshed.

Hot Module Replacement has a performance impact on compilation.

Also check the [webpack-dev-server documentation](http://webpack.github.io/docs/webpack-dev-server.html).

### Production compilation and server

``` text
# build the client bundle and the prerendering bundle
# and then start an nginx server to serve it
docker-compose -f docker-compose/production.yml up

open http://dockerhost:8000
```

The configuration is `webpack-production.config.js`.

The server is at `lib/server.js`

The production setting builds two configurations: one for the client (`build/public`) and one for the serverside prerendering (`build/prerender`).

#### Git Hooks
In order to commit, run:

```
npm install npm-sort -g
```
