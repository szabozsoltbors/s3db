ENVS := dev prod
PROFILE := aws_profile

check-env:
ifndef ENV
	$(error ENV must be provided!)
else
ifeq ($(filter $(ENV),$(ENVS)),)
	$(error ENV '$(ENV)' is not valid!)
endif
endif

check-region:
ifndef REGION
	$(error REGION must be provided!)
endif

init: check-env check-region
init: export AWS_PROFILE=$(PROFILE)
init: export TF_DATA_DIR=.terraform-$(ENV)-$(REGION)
init:
	tofu init \
		-reconfigure \
		-backend-config="key=backend/$(ENV)/$(REGION)/terraform.tfstate" \
		-backend-config=./backends/$(ENV)/$(REGION).tfvars

plan: check-env check-region
plan: export AWS_PROFILE=$(PROFILE)
plan: export TF_DATA_DIR=.terraform-$(ENV)-$(REGION)
plan:
	tofu plan \
		-var-file=environments/$(ENV)/$(REGION).tfvars \
		-out=tfplan-$(ENV)-$(REGION)

apply: check-env check-region
apply: export AWS_PROFILE=$(PROFILE)
apply: export TF_DATA_DIR=.terraform-$(ENV)-$(REGION)
apply:
	tofu apply -auto-approve tfplan-$(ENV)-$(REGION)
	rm -f tfplan-$(ENV)-$(REGION)

destroy: check-env check-region
destroy: export AWS_PROFILE=$(PROFILE)
destroy: export TF_DATA_DIR=.terraform-$(ENV)-$(REGION)
destroy:
	tofu destroy \
		-var-file=environments/$(ENV)/$(REGION).tfvars
