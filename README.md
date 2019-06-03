#Rancher plugin go version

This plugin is for solving JavaScript plugin which doesn't failed event if the drone workload isn't updated.

##How to test

```
export PLUGIN_PROJECT_API = ${Your Rancher workload API}
export PLUGIN_DEPLOY_IMAGE = ${Your docker image url}
export PLUGIN_ACCESS_KEY = ${Your Rancher API access key}
export PLUGIN_SECRET_KEY = ${Your Rancher API secret key}

go main.go
```

##HOW to use

You can use this way:

pipeline:
deploy-on-rancher:
    image: hazel910159/drone-deploy-plugin-go:v8
    project_api: ${Your Rancher workload API}
    access_key: ${Your Rancher API access key}
    secret_key: ${Your Rancher API secret key}
    deploy_image: ${Your docker image url}
    when:
      event: [push]
      branch: [beta, master]
