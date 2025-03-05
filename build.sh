docker build \
--build-arg AUTH_URL=$AUTH_URL \
--build-arg CLIENT_ID=$CLIENT_ID \
--build-arg CLIENT_SECRET=$CLIENT_SECRET \
--build-arg REDIRECT_URL=$REDIRECT_URL \
--build-arg ACCREF_TOKEN_URL=$ACCREF_TOKEN_URL \
-t fitness-summary --no-cache .
