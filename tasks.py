from pathlib import Path

from invoke import task  # pylint: disable=E0401

GIT_ROOT = Path(__file__).resolve().parent
ENV_FILE = Path(GIT_ROOT / ".env")
ONAIR_SH = Path.home() / "onair"
OFFAIR_SH = Path.home() / "offair"
BASH_PROFILE = Path.home() / ".bash_profile"
START_ONAIR_BLOCK = "### START ON AIR ###"
END_ONAIR_BLOCK = "### END ON AIR ###"

PRINT_LN = "-------------------"


def ensure_dot_env(context):
    print(PRINT_LN)
    print("👀 Checking '.env'...")

    if ENV_FILE.exists():
        print("⏭️  '.env' already exists, skipping.")
    else:
        print("✨ Creating '.env'...")
        context.run(f"cp {GIT_ROOT / '.env.example'} {ENV_FILE}")
        print("✅ '.env' created!")
    print(PRINT_LN)


@task
def say_hello(context):
    context.run("echo 'Hello World!'")


@task
def install(context):
    print(PRINT_LN)
    print("💾 Installing ON AIR...")
    deploy_scripts(context)
    write_bash_aliases(context)
    print("✅ ON AIR installed successfully!")
    print(PRINT_LN)


@task
def uninstall(context):
    """
    Removes all scripts, bash aliases, and virtrual environments
    """
    while True:
        confirm = input("🎙️ 🚨 Are you sure you wish to uninstall ON AIR? (y/n) ")
        first_letter_lowercase = confirm[0].lower()
        if confirm == "" or first_letter_lowercase not in ["y", "n"]:
            print("Please answer with yes or no!")
        else:
            break

    print(f"\n{PRINT_LN}")
    if first_letter_lowercase == "y":
        print("🗑️  Uninstalling ON AIR...")
        # remove scripts
        ONAIR_SH.unlink(missing_ok=True)
        OFFAIR_SH.unlink(missing_ok=True)
        # remove bash aliases
        context.run(
            f"sed -i.uninstall.onair.bak '/{START_ONAIR_BLOCK}/,/{END_ONAIR_BLOCK}/d' { BASH_PROFILE }"
        )
        # clean up python etc.
        print("🔥 ON AIR uninstalled successfully!")
    if first_letter_lowercase == "n":
        print("❌ Uninstall cancelled.")
    print(PRINT_LN)


@task
def write_bash_aliases(context):
    """
    Writes onair/offair bash aliases
    """
    print(PRINT_LN)
    print(f"👀 Checking '{BASH_PROFILE}' for project aliases...")
    with open(BASH_PROFILE, "r", encoding="utf-8") as profile:
        content = profile.read()
        if START_ONAIR_BLOCK in content:
            print("⏭️  Project aliases already exist, skipping.")
        else:
            print(f"🏷️  Writing project aliases to '{BASH_PROFILE}'...")
            context.run(f"cp {BASH_PROFILE} {BASH_PROFILE}.install.onair.bak")
            context.run(
                f"""cat <<EOT >> { BASH_PROFILE }
{START_ONAIR_BLOCK}
alias onair="$HOME/onair"
alias offair="$HOME/offair"
{END_ONAIR_BLOCK}
EOT"""
            )
            print(f"✅ Project aliases written to '{BASH_PROFILE}' successfully!")
    print(PRINT_LN)


@task
def install_dependencies(context):
    """
    perform a 'poetry install' to install python packages
    """
    context.run("poetry install")

    ensure_dot_env(context)


@task
def update_dependencies(context):
    """
    perform a 'poetry update' to update python packages
    """
    context.run("poetry update")


@task
def deploy_scripts(context):
    """
    Deploys onair/offair scripts at the user's home dir.
    """
    print(PRINT_LN)
    print(f"👀 Checking '{ Path.home() }' for convenence scripts...")
    if not ONAIR_SH.exists() and not OFFAIR_SH.exists():
        print("📜 Writing convenence scripts...")
        with open(ONAIR_SH, "w", encoding="utf-8") as on_file:
            onair_content = [
                "#!/bin/bash\n\n",
                f"cd {GIT_ROOT} || exit;\n\n",
                "./run-on-air.sh\n",
            ]
            on_file.writelines(onair_content)

        context.run(f"chmod +x {ONAIR_SH}")

        with open(OFFAIR_SH, "w", encoding="utf-8") as off_file:
            offair_content = [
                "#!/bin/bash\n\n",
                f"cd {GIT_ROOT} || exit;\n\n",
                "./run-off-air.sh\n",
            ]
            off_file.writelines(offair_content)

        context.run(f"chmod +x {OFFAIR_SH}")
        print("✅ Conveneince scripts written successfully!")
    else:
        print("⏭️  Conveneince scripts already installed, skipping.")


@task
def discover_process_names(context, query=None):
    """
    List running processes.
    """
    with context.cd(GIT_ROOT):
        if query:
            context.run(
                f"poetry run python host/discover_processes.py | grep -i {query}"
            )
        else:
            context.run("poetry run python host/discover_processes.py")


@task
def get_camera_state(context):
    """
    Gets the current camera state and writes it to the local stafefile.
    """
    with context.cd(GIT_ROOT):
        context.run("poetry run python host/get_vda_power_state.py")


@task
def get_device_state(context):
    """
    Gets the current homekit device state and writes it to the local stafefile.
    """
    with context.cd(GIT_ROOT):
        context.run("poetry run python host/get_device_state.py")


# @task
# def on_air(context):
#     """
#     Turns on the homekit device
#     """
#     with context.cd(GIT_ROOT):
#         context.run("poetry run python on_air.py")


# @task
# def off_air(context):
#     """
#     Turns off the homekit device
#     """
#     with context.cd(GIT_ROOT):
#         context.run("poetry run python off_air.py")


@task
def manage_cron(
    context,
    action: str = None,
    interval_min: int = 0,
    start_hour: int = 0,
    end_hour: int = 0,
    line_num: int = 0,
):
    # pylint: disable=[R0913,R0912,R1705]
    """
    manages the crontab. Requires an 'action'
    """
    with context.cd(GIT_ROOT / "scripts"):
        if action == "add":
            # TODO: check if onair has already been written to crontab.
            # ensure all options are set
            if not all([interval_min, start_hour, end_hour]):
                print(
                    "❌ Error: Must provide 'interval-min', 'start-hour', and 'end-hour' with '--action add'"
                )
                return
            else:
                # run checks on hour values
                if start_hour > end_hour:
                    print(
                        f"❌ Error: 'start-hour' ({start_hour}) cannot be greater than 'end-hour' ({end_hour}).\n   Ensure you are using 24H time."
                    )
                    return
                elif start_hour == end_hour:
                    cron_hours = end_hour
                else:
                    cron_hours = f"{start_hour}-{end_hour}"
                # run checks on minute values
                if interval_min == 0:
                    cron_min = "*"
                else:
                    cron_min = f"*/{interval_min}"
                # run the script
                context.run(
                    f"./cron-mgmt.sh add '{cron_min} {cron_hours} * * 1-5 {GIT_ROOT}/run-app-status-light.sh > /dev/null 2>&1'"
                )
                print("🎙️ ON AIR entry written to crontab")
        elif action == "list":
            context.run("./cron-mgmt.sh list")
        elif action == "remove":
            if not line_num:
                print("❌ Error: Must provide 'line-num' with '--action remove'")
            else:
                context.run(f"./cron-mgmt.sh remove {line_num}")
                print(f"✅ line {line_num} removed from crontab.")
        else:
            print("❌ Error: Must provide an 'action'")
