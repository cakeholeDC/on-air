from host.actions import get_and_log_device_status

if __name__ == "__main__":
    homekit_state = get_and_log_device_status()
    print(homekit_state)
