# go-rest-api
Go REST API proj

## 🚀 Getting Started

### Prerequisites
* **Go** (1.23 or higher)
* **OpenSSL** (Required for generating local SSL certificates)

### Local SSL Setup (HTTP/2)
To support **HTTP/2**, this server requires a secure connection via TLS. The `cert.pem` and `key.pem` files are excluded from this repository for security. 

To generate your own local certificates for development, run the following command in the project root:

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config openssl.cnf
