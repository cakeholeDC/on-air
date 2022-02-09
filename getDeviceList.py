from LightControl import loop, listDevices

if __name__ == '__main__':
    loop.run_until_complete( listDevices() )