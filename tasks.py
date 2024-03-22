from invoke import task
from pathlib import Path

GIT_ROOT = Path(__file__).resolve().parent

@task
def say_hello(context):
    context.run("echo 'Hello World!'")

@task
def install_dependencies(context):
    """
    perform a 'poetry install' to install python packages
    """
    context.run("poetry install --no-root")

@task
def update_dependencies(context):
    """
    perform a 'poetry update' to update python packages
    """
    context.run("poetry update")
@task
def manage_cron(context, action: str=None, interval_min: int=0, start_hour: int=0, end_hour: int=0, line_num: int=0):
    """
    manages the crontab. Requires an 'action'
    """
    with context.cd(GIT_ROOT / "scripts"):
        if action == "add":
            # TODO: check if onair has already been written to crontab.
            # ensure all options are set
            if not all([interval_min, start_hour, end_hour]):
                context.run(f"echo '❌ Error: Must provide `interval-min`, `start-hour`, and `end-hour` with `--action add`'")
                return
            else:
                # run checks on hour values
                if start_hour > end_hour:
                    context.run(f"""
                        echo '❌ Error: `start-hour` ({start_hour}) cannot be greater than `end-hour` ({end_hour}).'
                        echo '   Ensure you are using 24H time.'
                    """)
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
                # TODO: update the cronjob.
                context.run(f"./cron-mgmt.sh add '{cron_min} {cron_hours} * * 1-5 {GIT_ROOT}/run-app-status-light.sh > /dev/null 2>&1'")
                context.run("echo '🎙️ on-air entry written to crontab.'")
        elif action == "list":
            context.run("./cron-mgmt.sh list")
        elif action == "remove":
            if not line_num:
                context.run("echo 'fail'")
            else:
                context.run(f"./cron-mgmt.sh remove {line_num}")
                context.run(f"echo '✅ line {line_num} removed from crontab.'")
        else:
            context.run("echo 'must provide an action'")

