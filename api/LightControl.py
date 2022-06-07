import aiohttp
import os
import pysmartthings
from api.Light import Light
from dotenv import load_dotenv
from api.colors import RED

load_dotenv()
token = os.getenv('SMARTTHINGS_TOKEN')
device_name = os.getenv('DEVICE_NAME')
lights_object = Light()

async def lights_on(object=lights_object, name=device_name):
    '''
    turns on device_name of type=object
    '''
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await find_devices(api, object, name)
        print(f"Turning on: {RED}{name}")
        await object.on()


async def lights_off(object=lights_object, name=device_name):
    '''
    turns off device_name of type=object
    '''
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await find_devices(api, object, name)
        print(f"Turning off: {RED}{name}")
        await object.off()


async def get_device_status(name=device_name):
    print(f"getting api.device.status for {name}")
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        devices = await api.devices()
        for device in devices:
            if device.label == name:
                await device.status.refresh()
                print(f"{name} = {device.status.switch}")
                return device.status.switch


async def find_devices(api, object=lights_object, name=device_name):
    devices = await api.devices()

    for device in devices:
        if device.label == name:
            object.device = device


async def get_locations(api):
    locations = await api.locations()

    print_large_divider()
    print(f'{len(locations)} Locations')
    for location in locations:
        print_small_divider()
        print(f'Name: {location.name}')
        print(f'\tId: {location.location_id}')


async def list_devices():
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await getDevices(api)
        print_large_divider()


async def getDevices(api):
    '''
    prints a list of api.devices to the console
    '''
    devices = await api.devices()

    print_large_divider()
    print(f'{len(devices)} Devices')
    for device in devices:
        print_small_divider()
        await device.status.refresh()
        print(f'Label: {device.label}')
        print(f'\t Type: {device.name}')
        print(f'\tId: {device.device_id}')
        print(f'\tCapabilities: {device.capabilities}')
        print(f'\tStatus: {device.status.switch}')


def print_large_divider():
    print('---------------------------------------------------------------------------------------------------------')


def print_small_divider():
    print('-----------------------------------------------------------------------')
