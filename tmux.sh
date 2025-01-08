#!/bin/sh

tmux new-session -d -s bugbase -n bugbase

tmux new-window -c ./app-issue -n services
tmux send-keys "make dev" C-m
tmux split-window -c ./app-static
tmux send-keys "make dev" C-m
tmux split-window -c ./database -h
tmux send-keys "ls" C-m
