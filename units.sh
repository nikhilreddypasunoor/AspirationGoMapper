#!/bin/bash
go test -v `(go list ./...)` -coverprofile=coverage.out