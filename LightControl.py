import aiohttp
import os
import asyncio
import pysmartthings
from random import randint
from time import sleep
from Light import Light
from dotenv import load_dotenv

load_dotenv()
token = os.getenv('TOKEN')
device_name = os.getenv('DEVICE_NAME')
loop = asyncio.get_event_loop()
lights_object = Light()

async def lightsOn(object, name):
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await findDevices(api, object, name)
        print(f"Turning on {name}...")
        await object.on()


async def lightsOff(object, name):
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await findDevices(api, object, name)
        print(f"Turning off {name}...")
        await object.off()


async def listDevices():
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        # * Prints all devices
        # await getLocations(api) # CURRENT API KEY DOES NOT ALLOW LOCATIONS
        await getDevices(api)
        largeDivider()


async def getDeviceStatus(name):
    print(f"getting device status for {name}")
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        devices = await api.devices()
        for device in devices:
            if device.label == name:
                await device.status.refresh()
                print(f"{name} = {device.status.switch}")
                return device.status.switch


async def findDevices(api, object, name):
    devices = await api.devices()

    for device in devices:
        if device.label == name:
            object.device = device


async def getLocations(api):
    locations = await api.locations()

    largeDivider()
    print(f'{len(locations)} Locations')
    for location in locations:
        smallDivider()
        print(f'Name: {location.name}')
        print(f'\tId: {location.location_id}')


async def getDevices(api):
    devices = await api.devices()

    largeDivider()
    print(f'{len(devices)} Devices')
    for device in devices:
        smallDivider()
        await device.status.refresh()
        print(f'Label: {device.label}')
        print(f'\t Type: {device.name}')
        print(f'\tId: {device.device_id}')
        print(f'\tCapabilities: {device.capabilities}')
        print(f'\tStatus: {device.status.switch}')


def largeDivider():
    print('---------------------------------------------------------------------------------------------------------')


def smallDivider():
    print('-----------------------------------------------------------------------')


def weirdRandomLights(object, name, low_num = 0, high_num = 60):
    subtract_from_high = 0
    while True:
        try:
            num = randint(low_num, high_num - int(subtract_from_high))

            if num == low_num:
                loop.run_until_complete(lightsOff(object, name))
                subtract_from_high=0
                print(f'Turning off')
            elif num == low_num+1:
                loop.run_until_complete(lightsOn(object, name))
                subtract_from_high=0
                print(f'Turning on')
            else:
                print(num)
                subtract_from_high+=0.4
                
            sleep(1)
            print(f'Max: {high_num - int(subtract_from_high)}')
        except KeyboardInterrupt:
            exit()

#* How to find device names
# loop.run_until_complete( listDevices() )

# weirdRandomLights(lights_object, device_name)

#* syntax to run async function
# loop.run_until_complete( lightsOn(lights_object, device_name) )
# loop.run_until_complete( lightsOff(lights_object, device_name) )

#! For some reason logging doesnt work in async functions or something idk because this is my first time using async functions
#TODO Make async functions work with logging