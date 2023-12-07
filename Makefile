include .env


login-gcloud:
	@gcloud auth login

setup-gcloud:
	gcloud config set project ${GCP_PROJECT_ID}
	gcloud config set compute/region ${GCP_REGION}
	gcloud config set compute/zone ${GCP_ZONE}

# Cloud RUN に対して API サーバーをデプロイする
deploy-api:
	@gcloud builds submit \
	--config ./cloudbuild.api.yml \
	--substitutions=\
_DOCKER_FILE=Dockerfile.prod,\
_DOCKER_STAGE_NAME=api-runner,\
_PROJECT_ID=${GCP_PROJECT_ID},\
_PROJECT_NAME=${GCP_API_PROJECT_NAME},\
_REGION=${GCP_REGION},\
_IMAGE=${GCP_API_IMAGE}

# # Cloud SQL に対してマイグレーションを実行する
# deploy-sql:
# 	@gcloud builds submit \
# 	--config ./cloudbuild.sql.yml \
# 	--substitutions=\
# _PROJECT_ID=${GCP_PROJECT_ID},\
# _REGION=${GCP_REGION},\
# _INSTANCE=${GCP_SQL_INSTANCE},\
# _URL=${GCP_SQL_URL}

deploy-sql:
	atlas migrate apply --url ${DB_CONNECTION_URL} --dir file://migrations
