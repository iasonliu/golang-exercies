#See Makefile

## Preparation

### Terraform
Download https://releases.hashicorp.com/terraform/0.12.26/ , add to path

### Ansible

```
brew install ansible
/usr/local/opt/ansible/libexec/bin/pip install pywinrm
/usr/local/opt/ansible/libexec/bin/pip install "ansible[azure]"
```

### az cli

```
brew install az
az login
```

## Run Sandbox

```
make init_env # ask for ~/.azure/credentials file if using other's subscription
make tf_init
make tf_plan
make tf_apply
make ansible_list
```