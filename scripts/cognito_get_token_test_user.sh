#!/bin/bash
set -euo pipefail

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

# chmod +x scripts/cognito_get_token_test_user.sh
# USER_POOL_ID= USER_NAME= NEW_PASSWORD= CLIENT_ID= ./scripts/cognito_get_token_test_user.sh
