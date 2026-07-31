>[!Note]
> This project is in early phase of development. Please handle with care.

# s3db

The main goal of the project is to build a simple cheap document database on top of AWS S3 and Lambda Services.

#### Pre-requirements

1. [GoLang](https://go.dev)
2. [AWS CLI](https://aws.amazon.com/cli)
3. [OpenTofu](https://search.opentofu.org/provider/opentofu/aws/latest)


#### Build and run the project

1. Execute the steps of the `./worker/README.md` file to compile the GoLang project.
2. Execute the steps of the `./infrastructure/src/README.md` file to create the needed infrastructure.
