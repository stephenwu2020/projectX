#!/bin/bash
# https://www.stuartellis.name/articles/shell-scripting/#enabling-better-error-handling-with-set
set -Eeuo pipefail
 
# Based on mongo/docker-entrypoint.sh
# https://github.com/docker-library/mongo/blob/master/docker-entrypoint.sh#L303
if [ "$MONGO_INITDB_ROOT_USERNAME" ] && [ "$MONGO_INITDB_ROOT_PASSWORD" ]; then
			rootAuthDatabase='admin'

			"${mongo[@]}" "$rootAuthDatabase" <<-EOJS
				db.createUser({
					user: $(_js_escape "$MONGO_INITDB_ROOT_USERNAME"),
					pwd: $(_js_escape "$MONGO_INITDB_ROOT_PASSWORD"),
					roles: [ { role: 'root', db: $(_js_escape "$rootAuthDatabase") } ]
				})
			EOJS
fi