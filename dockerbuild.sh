#!/bin/bash
docker build -t gomysqlapi .
docker run -it --rm --name my-running-app -p 30801:30801 gomysqlapi