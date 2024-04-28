#!/bin/bash

cmd=${1:-up}
project_name=""
file_prefix=""

execute-docker-compose () {
  opts="-f ${file_prefix}docker-compose.yml"

  if [ -n "$project_name" ]; then
    opts="$opts -p $project_name"
  fi

  docker-compose $opts $@
}

stop-docker-compose () {
  execute-docker-compose stop
}

if [ $cmd = 'up' ] && [ $# -le 1 ]; then
  execute-docker-compose up
  stop-docker-compose

elif [ $cmd = 'stop' ]; then
  stop-docker-compose

elif [ $cmd = 'bash' ]; then
  execute-docker-compose exec -w /backend memoria-api /bin/bash

elif [ $cmd = 'bash-api' ]; then
  execute-docker-compose exec -w /backend/services/memoria-api memoria-api /bin/bash

elif [ $cmd = 'bash-db' ]; then
  execute-docker-compose exec memoria-db /bin/bash

else
  execute-docker-compose $@
fi
