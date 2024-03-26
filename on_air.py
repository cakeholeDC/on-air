from host.actions import turn_on_and_log_status
from host.colors import RED, GREEN, COLOR_NONE

if __name__ == "__main__":
    turn_on_and_log_status()
    print(f"{GREEN}------------------")
    print(f"{GREEN}| --- {RED}ON AIR {GREEN}--- |")
    print(f"{GREEN}------------------{COLOR_NONE}")
    # os.system("printenv")
