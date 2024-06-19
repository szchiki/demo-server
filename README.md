# Demo server

## Running the server
You can run the server locally using the container. Make sure you're in project root.
```bash
  docker build -t demo-server -f cmd/Dockerfile . && docker run -it demo-server
```

Deploying to cloud run:
```bash
  gcloud builds submit --config cmd/cloudbuild.yaml --region europe-west1 .
```

## Making requests
For local servers, call port :8080. Can be configured with PORT env variable.

`curl 'http://localhost:8080/delivery?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'`

For cloud run servers, call https://delivery-demo-myzvsyegga-lz.a.run.app, like so:

`curl 'https://delivery-demo-myzvsyegga-lz.a.run.app/delivery?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'`

## Structure
```
/
    /cmd -- Contains our main package, multiple if you want more than one application.
        Dockerfile -- Builds our container.
        cloudbuild.yaml -- Configuration for our cloud build pipe. Pretty much just runs docker build, push and then deploys from a cloud machine.
        /transport -- Contains our transport layer code, most often the HTTP/RPC server.
    /internal -- Contains base structures and integrations we'll need. Can be replaced by pkg if you prefer an open approach. A database of route legs could be stored here.
    /services -- Contains one or more of our services, in this case juse the one. Here's the business logic specific to our service/product.
cloud_setup.sh -- A living file containing steps required to recreate our cloud environment. A very simple pre-cursor to a full scale service like terraform.
```