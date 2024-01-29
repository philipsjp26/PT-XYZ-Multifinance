# PT-XYZ Finance

## Description
PT XYZ Multifinance adalah salah stau perusahaan pembiayaan motor dan mobil di indonesia,
yang mempunyai komitmen menyediakan solusi pembiayaan terhadap masyarakat melalui teknologi 
untuk meningkatkan kualitas hidup masyarakat umum.

Berdasarkan background multifinance, kita akan membuat server kecil yang membutuhkan environment yang __scalable, maintanable, testable__ dll.

## Installation
Use the makefile syntax to install dependencies
```shell
make clean && make install
```

After install, set up configuration file on app.yaml (conf/app.yaml)

## Usage
Run migration using __goose__ as migration tools. <br />

__note__: __Please create database first !__
```shell
go run main.go db:migrate up
```

Build docker images
```shell
docker build -t {image_name} -f deployment/Dockerfile .
```
## REST API:
`POST :8000/api/v1/customer`
```curl --location ':8000/api/v1/customer' \
--header 'Content-Type: application/json' \
--data '{
    "identity_number": "",
    "full_name": "",
    "legal_name": "",
    "place_of_birth": "",
    "date_of_birth": "",
    "salary": 0
}'
```

`POST :8000/api/v1/customer/limit`
```
curl --location ':8000/api/v1/customer/limit' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id": 1,
    "tenor": 1,
    "limit_amount": 1000000
}'
```

`POST :8000/api/v1/customer/transaction`
```
curl --location ':8000/api/v1/customer/transaction' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id": 1,
    "otr": 6000000,
    "admin_fee": 25000,
    "installment_amount": 650000,
    "interest_amount": 10000
}'
```
For handling concurrency on this service, we implement some method on endpoint `:8000/api/v1/customer/transaction/:contract_number` :

```
curl --location --request PUT ':8000/api/v1/customer/transaction/018d545a-f120-7d9b-816e-5f23eba70da2' \
--header 'Content-Type: application/json' \
--data '{
    "admin_fee": 2000
}'
```
- Optimistic Locking
- Using mutex

## System Architecture
Its using hexagonal architecture as base structure

![alt text ](https://github.com/philipsjp26/PT-XYZ-Multifinance/blob/master/diagram.png?raw=true)

<br />

Example common architecture

![alt text ](https://github.com/philipsjp26/PT-XYZ-Multifinance/blob/master/architecture.png?raw=true)

<br />
Database Diagram

![alt text ](https://github.com/philipsjp26/PT-XYZ-Multifinance/blob/master/database.png?raw=true)