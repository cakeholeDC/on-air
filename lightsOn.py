import os
from LightControl import loop, lightsOn, lights_object, device_name
from hostActions import write_log_file

loop.run_until_complete( lightsOn(lights_object, device_name) )

write_log_file(True)
