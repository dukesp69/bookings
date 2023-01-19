#!/bin/bash


#build project && and if succceded run binary
go build -o bookings cmd/web/*.go && ./bookings