build:
	@go build -o onair .
	@go build -o /usr/local/bin/onair .

launchd-interval:
	@secs=$$(date +%S); sleep $$((60 - $$secs))
	@$(MAKE) launchd-load

launchd-load:
	@cp ./local.onair.plist $(HOME)/Library/LaunchAgents
	@launchctl load $(HOME)/Library/LaunchAgents/local.onair.plist

launchd-unload:
	@launchctl unload $(HOME)/Library/LaunchAgents/local.onair.plist
	@rm $(HOME)/Library/LaunchAgents/local.onair.plist

launchd-status:
	@launchctl list | grep local.onair || echo "local.onair.plist not loaded"
