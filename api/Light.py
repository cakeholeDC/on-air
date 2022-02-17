class Light:
    device = None
    switch = None
    switch_level = None

    async def on(self):
        result = await self.device.switch_on()
        assert result == True
        self.switch = True
        return result

    async def off(self):
        result = await self.device.switch_off()
        assert result == True
        self.switch = False
        return result

    async def setLevel(level, rate, self):
        result = await self.set_level([level, rate])
        assert result == True
        self.switch_level = level
        return result