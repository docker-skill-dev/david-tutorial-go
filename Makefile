configure:
	docker run --init -it --rm --pull=always \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v $${PWD}:/skill \
		atomist/local-skill --pwd /skill --workspace $${ATOMIST_WORKSPACE} --apikey $${ATOMIST_API_KEY} --host-dir $${PWD} --watch

run:
	go run "." 2>> debug.log

logs:
	tail -f debug.log

.PHONY: part1 part2 part3 part4

part1:
	rm bugs.go closed_bug.go webhooks.go datalog/subscription/on_bug_close.edn datalog/schema/bugs.edn || true
	cp -r tutorial/part1/* ./

part2: part1
	cp -r tutorial/part2/* ./

part3: part2
	cp -r tutorial/part3/* ./

part4: part3
	cp -r tutorial/part4/* ./