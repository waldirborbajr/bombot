#!/bin/bash

SESSION_NAME="development"

# Create a new Tmux session
tmux -u new-session -d -s $SESSION_NAME

#This script requires autojump to be installed and alias added as 'j'

# Create windows with specified names and execute commands in each window
tmux -u new-window -t $SESSION_NAME:1 -n "hot"
tmux send-keys -t $SESSION_NAME:1 'nvim main.go' C-m

tmux -u new-window -t $SESSION_NAME:2 -n "ngrok"
tmux send-keys -t $SESSION_NAME:2 'ngrok http 9090' C-m

tmux -u new-window -t $SESSION_NAME:3 -n "start bot"
tmux send-keys -t $SESSION_NAME:3 'sleep 5 && ./botON-OFF on' C-m

tmux -u new-window -t $SESSION_NAME:4 -n "Makefile run"
tmux send-keys -t $SESSION_NAME:4 'sleep 5 && make run' C-m

tmux -u new-window -t $SESSION_NAME:5 -n "Lazygit"
tmux send-keys -t $SESSION_NAME:5 'lazygit' C-m

# Attach to the newly created session
tmux -u attach-session -t $SESSION_NAME