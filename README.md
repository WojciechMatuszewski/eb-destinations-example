# EB Connections

This repository contains a SAM template showing how one might use the newly introduced _Event Bridge API Destinations_.

## Deployment

Deploying the infrastructure in this repository is a two phase process (see _Learnings_ for more detail)

1. Comment out everything after `# === EB Stuff === #`.
2. Deploy the stack
3. Change the _Secrets Manager_ secret value to the actual _Cognito secret_
4. Uncomment the lines commented in step 1
5. Deploy the infrastructure again

## Learnings

- Some resources do not support _SSM secure string parameters_. You need to use _Secrets Manager_

- the ApiDestination will try to authorize the request during the deployment. This means that you will probably have to do 2 phase deployment. In our case, we first need to deploy the _Cognito_ related infrastructure, change the _Secrets Manager_ secret, then deploy the _API Destination_ and the _EB Rule_

- _Cognito_ exposes a nice way to perform _machine-to-machine_ communication. I finally understand what is the _generate secret_ setting for!
