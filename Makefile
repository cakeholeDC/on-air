build:
	@go build -ldflags="-linkmode=external" -o onair .
	@mkdir -p $(HOME)/.local/bin
	@cp ./onair $(HOME)/.local/bin/onair

launchd-interval:
	@secs=$$(date +%S); sleep $$((60 - $$secs))
	@$(MAKE) launchd-load

launchd-load:
	@sed -i '' 's|!HOME!|$(HOME)|' ./local.onair.plist
	@sed -i '' 's|!USER!|$(USER)|' ./local.onair.plist
	@cp ./local.onair.plist $(HOME)/Library/LaunchAgents
	@launchctl load $(HOME)/Library/LaunchAgents/local.onair.plist

launchd-unload:
	@launchctl unload $(HOME)/Library/LaunchAgents/local.onair.plist
	@rm $(HOME)/Library/LaunchAgents/local.onair.plist

launchd-reload: launchd-unload launchd-load

launchd-status:
	@launchctl list | grep local.onair || echo "local.onair.plist not loaded"
