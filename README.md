<h1>Rancher plugin go version</h1>

This plugin is for solving JavaScript plugin which doesn't failed event if the drone workload isn't updated.

<h2>How to test</h2>

```
export PLUGIN_PROJECT_API = ${Your Rancher workload API}
export PLUGIN_DEPLOY_IMAGE = ${Your docker image url}
export PLUGIN_ACCESS_KEY = ${Your Rancher API access key}
export PLUGIN_SECRET_KEY = ${Your Rancher API secret key}

go main.go
```

<h2>HOW to use</h2>

You can use this way:

```
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
```
