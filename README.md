# Example

## Get started

### Preparing
- Install Docker - https://www.docker.com/community-edition
- Install Docker Compose - https://docs.docker.com/compose/install/
- Install dependencies - `make install`
- Run the application - `make run`

### Request example
`curl -XPOST -H "Content-type: application/json" -d '{"id":1,"name":"User Name","email":"user@gmail.com"}' -A 'Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5' 'localhost:8000'`

### Useful commands
- `make install` - installing dependencies and configuring of the environment
- `make run` - run the server
- `make build` - build application
- `make vendor` - update vendor
- `make get pkg=github.com/usr/pkg` - install package
- `make bash` - enter into golang container
- `make up` - start docker
- `make down` - strop docker
- `make ps` - show docker status
