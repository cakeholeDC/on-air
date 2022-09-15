import os

LIGHT_STATUS = os.getenv('LIGHT_STATUS')
def write_light_status(status):
    '''
    Writes `True` or `False` to the local log file
    '''
    log_file = open(LIGHT_STATUS, "w") # pylint: disable=[W1514, R1732]
    log_file.write(str(status))
    log_file.write("\n")
    log_file.close()


CAMERA_STATUS = os.getenv('CAMERA_STATUS')
def write_camera_status(status):
    '''
    Writes `True` or `False` to the local log file
    '''
    log_file = open(CAMERA_STATUS, "w") # pylint: disable=[W1514, R1732]
    log_file.write(str(status))
    log_file.write("\n")
    log_file.close()
    