class Light:
    device = None
    switch = None
    switch_level = None

    async def on(self): # pylint: disable=[C0103]
        result = await self.device.switch_on()
        assert result is True
        self.switch = True
        return result

    async def off(self): # pylint: disable=[C0103]
        result = await self.device.switch_off()
        assert result is True
        self.switch = False
        return result

    async def setLevel(level, rate, self): # pylint: disable=[C0103]
        result = await self.set_level([level, rate])
        assert result is True
        self.switch_level = level
        return result
