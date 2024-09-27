# K8s Workshop

`git clone git@github.com:apoorv-deriv/k8s-ctf.git`

## Connect to EKS cluster.

Steps:

- [ ] Clone https://github.com/deriv-iac/iac-eks
- [ ] Run `make ssm-qa-admin`
- [ ] Open https://device.sso.eu-west-1.amazonaws.com/ and login.
- [ ] And enter the code displayed at the bottom.

Try `kubectl get ns` to check if your connected successfully

## Deploy a test application

We will deploy the application the way developers do it, using https://github.com/regentmarkets/environment-manifests-qa/

Steps:

- [ ] Clone https://github.com/regentmarkets/environment-manifests-qa/. And switch to your QA branch
- [ ] Create a `mkdir -p k8s` folder in repo.
- [ ] Now copy the contents of `k8s/test-app` from this repo to `k8s`. `cp -r ../k8s-ctf/k8s/test-app k8s/`
- [ ] 

