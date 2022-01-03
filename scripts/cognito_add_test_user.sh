#!/bin/bash
set -euo pipefail

# create user
aws cognito-idp admin-create-user \
  --user-pool-id $USER_POOL_ID \
  --username $USER_NAME \
  --temporary-password $TEMPORARY_PASSWORD \
  --user-attributes Name=email,Value="$USER_NAME" Name=name,Value="$NAME" Name=family_name,Value="$FAMILY_NAME" \
  --message-action SUPPRESS

echo "user created"

# first authentication
SESSION=$(aws cognito-idp admin-initiate-auth \
  --user-pool-id $USER_POOL_ID \
  --client-id $CLIENT_ID \
  --auth-flow ADMIN_NO_SRP_AUTH \
  --auth-parameters USERNAME=$USER_NAME,PASSWORD=$TEMPORARY_PASSWORD \
  | grep -zoP '"Session":\s*"\K[^"]*')

echo "auth initialized"

# respond to the challenge
aws cognito-idp admin-respond-to-auth-challenge \
  --user-pool-id $USER_POOL_ID \
  --client-id $CLIENT_ID \
  --challenge-name NEW_PASSWORD_REQUIRED \
  --challenge-responses USERNAME=$USER_NAME,userAttributes.name=$NAME,NEW_PASSWORD=$NEW_PASSWORD \
  --session $SESSION

echo "password updated"

# get a token
TOKEN=$(aws cognito-idp admin-initiate-auth \
  --user-pool-id $USER_POOL_ID \
  --client-id $CLIENT_ID \
  --auth-flow ADMIN_NO_SRP_AUTH \
  --auth-parameters USERNAME=$USER_NAME,PASSWORD=$NEW_PASSWORD \
  | grep -zoP '"IdToken":\s*"\K[^"]*')

echo "token:"
echo ""
echo $TOKEN

# remember to set the env vars

# chmod +x scripts/cognito_add_test_user.sh
# USER_POOL_ID= USER_NAME= TEMPORARY_PASSWORD= NEW_PASSWORD= NAME= FAMILY_NAME="" CLIENT_ID= ./scripts/cognito_add_test_user.sh
