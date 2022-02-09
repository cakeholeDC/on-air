# This is python code to mess around with my lights

In order to use it, the first thing one needs to do is [get a smartthings personal access token](https://account.smartthings.com/tokens).
Secondly, one must copy the [.env.example](./.env.example) file into a [.env file](./.env).
Copy this token and put it in your .env file (example can be found [here](./.env.example).)

# Run it
To get it working, go into [LightControl.py](./LightControl.py), and run the program.
This will give an output of all devices and their states. Find your desired device and add it to your [.env file](./.env) as the value for `DEVICE_NAME`.

Comment the line `loop.run_until_complete( listDevices() )` and uncomment the line `#weirdRandomLights(my_lights_object, device_name)`

Run the program again and it should turn on and off the device specified in a semi-random manner.


# Terminal Aliases
Add the following to your `.bash_profile` to enable termainal aliases

```bash
# OnAir Aliases
alias onair="~/on-air"
alias offair="~/off-air"
```