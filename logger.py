import os

STATUS_LOG = os.getenv('STATUS_LOG')
def write_log_file(status):
    '''
    Writes `True` or `False` to the local log file
    '''
    log_file = open(STATUS_LOG, "w") # pylint: disable=[W1514, R1732]
    log_file.write(str(status))
    log_file.write("\n")
    log_file.close()
