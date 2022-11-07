db-deploy:
	@echo "Deploying database"

gcloud-deploy:
	gcloud builds submit --tag ${GCP_IMAGE_IMAGE} .
