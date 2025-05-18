#!/bin/bash

ENV_FILE="./data/secrets.env"

# Create or overwrite the .env file
touch $ENV_FILE

# Write environment variables to the .env file in one block
{
  echo "TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN"
} >> $ENV_FILE