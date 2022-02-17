import os

STATUS_LOG = os.getenv('STATUS_LOG')
def write_log_file(status):
    f = open(STATUS_LOG, "w")
    f.write(str(status))
    f.write("\n")
    f.close()
