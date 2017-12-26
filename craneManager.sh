#!/bin/bash

if [[ $1 == "start" ]]; then
  echo "./crane" | bash
  echo "Crane listener started"
elif [[ $1 == "stop" ]]; then
  echo "kill $(lsof -i :3333 | grep crane | awk '{print $2}')" | bash
  echo "Crane listener stopped"
elif [[ $1 == "status" ]]; then
  echo "lsof -i :3333" | bash
else
  echo "No command line arguments provided. Please provide a command, such as"
  echo "start - starts the crane listener on port 3333"
  echo "stop - stops the crane listener"
  echo "status - shows running processes on port 3333"
fi
