# Golang Docker Base

Golang Docker Base is an extremely simple Go & Docker webserver setup for quick testing locally as well as Google Cloud Run. The binary is built in an alpine container, and copied over to a scratch container, which ends up at around 6.42MB in size.

## Usage

1. Clone the repo
2. Build the container `docker build -t {name} .`
3. Run the container `docker run -p 8080:8080 {name}`