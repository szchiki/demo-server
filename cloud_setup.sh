#!/bin/bash

export PROJECT_ID="delivery-demo-426913"

gcloud config set project $PROJECT_ID

gcloud artifacts repositories create delivery-demo \
    --repository-format docker \
    --location europe-north1 \
    --description "Docker repository" \
    --project $PROJECT_ID | exit 0