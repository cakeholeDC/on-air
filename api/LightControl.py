import aiohttp
import os
import pysmartthings
from api.Light import Light
from dotenv import load_dotenv
from api.colors import RED

load_dotenv()
token = os.getenv('SMARTTHINGS_TOKEN')
device_id = os.getenv('DEVICE_GUID')
lights_object = Light()

async def lights_on(object=lights_object, guid=device_id):
    '''
    turns on device_name of type=object
    '''
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await get_device_by_guid(api, object, guid)
        print(f"Turning on: {RED}{object.device.label}")
        await object.on()

async def lights_off(object=lights_object, guid=device_id):
    '''
    turns off device_name of type=object
    '''
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        await get_device_by_guid(api, object, guid)
        print(f"Turning off: {RED}{object.device.label}")
        await object.off()

async def get_device_status(guid=device_id):
    print(f"getting api.device.status for {object.device.label}")
    async with aiohttp.ClientSession() as session:
        api = pysmartthings.SmartThings(session, token)

        devices = await api.devices()
        for device in devices:
            if device.device_id == guid:
                await device.status.refresh()
                print(f"{device.label} = {device.status.switch}")
                return device.status.switch

async def get_device_by_guid(api, object=lights_object, guid=device_id):
    devices = await api.devices()

    for device in devices:
        if device.device_id == guid:
            object.device = device
            break


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
    # TODO: add color to the output for easier human parsing
    devices = await api.devices()

    print_large_divider()
    print(f'{len(devices)} Devices')
    for device in devices:
        print_small_divider()
        await device.status.refresh()
        print(f'Name: {device.label}') # NOTE: device.label is the name set in the SmartThings App
        print(f'\tType: {device.name}')
        # NOTE: device.name is configurable via the SmartThings IDE >> https://stdavedemo.readthedocs.io/en/latest/ref-docs/device-ref.html#name
            # Personal config is Manufacturer:Type:uuid (ex Hue:Dimmable:bed-2)
            # TODO: `type` is believed to be an API value that we could use instead (device.type)
        print(f'\tId: {device.device_id}')
        print(f'\tCapabilities: {device.capabilities}')
        print(f'\tStatus: {device.status.switch}')


def print_large_divider():
    print('---------------------------------------------------------------------------------------------------------')


def print_small_divider():
    print('-----------------------------------------------------------------------')
