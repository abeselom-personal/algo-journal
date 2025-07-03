#!/bin/bash
msg="Update: $(date '+%Y-%m-%d')"
git add .
git commit -m "$msg"
git push
