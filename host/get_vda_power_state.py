from host.actions import get_and_log_vdc_status

if __name__ == "__main__":
    state = get_and_log_vdc_status()
    print(state)
