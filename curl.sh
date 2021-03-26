#!/usr/bin/env bash

curl $(pulumi stack output bucketEndpoint)
