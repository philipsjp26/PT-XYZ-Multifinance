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

## System Architecture
Its using hexagonal architecture as base structure

![alt text ](https://github.com/philipsjp26/PT-XYZ-Multifinance/blob/master/diagram.png?raw=true)

Example common architecture
