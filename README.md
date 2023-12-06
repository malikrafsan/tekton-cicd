# Tekton CI/CD

## Summary
> This project is a demo project for tekton as CI/CD configuration on Kubernetes

## How to Run

### Generate Keys
1. Generate SSH Key
```sh
ssh-keygen -t rsa -b 4096 -C "tekton@tekton.dev"  # store at ~/.ssh/tekton_rsa

cat ~/.ssh/tekton_rsa | base64   # put the result into tekton/secrets/git-ssh.yaml on `ssh-privatekey: <base64 encoded private key>`
```

2. Generate Image Registry Config
    
    a. Fill the `config.json` template
    
    b. Generate config on base64
    ```sh
    cat config.json | base64   # put the result into tekton/secrets/docker-credential.yaml
    ```

### Apply Tekton
```sh
chmod +x tekton-apply
./tekton-apply
```

### Run Tekton Pipeline
```sh
chmod +x tekton-run
./tekton-run
```
