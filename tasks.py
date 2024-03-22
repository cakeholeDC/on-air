from invoke import task

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
