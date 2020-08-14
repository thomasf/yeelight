# yeelight


I have been keeping myself from pushing it to github because I want to clean
the code and add documentation but I havent had any time for it after I got a
working mqtt bridge.

I have been running this in production at home with about 10 bulbs for about 3
year so maybe someone else has use of it even if the code isn't pretty.

The only problem I have notices is that the reconnection code some times doesnt
work right after a bulb is restarted by disconnecting it's power, other than
that I don't know about any issues.

When I get back to working on this the Go API and MQTT stuff might start
changing....


# MQTT overview (TODO: NOT COMPLETE)

## topics

There are three base command topics depending on the type of device.

### commands


Sending commands to control lights:

- `ylcolor/command/[devicename]/[commandname]`
- `ylmono/command/[evicename]/[commandname]`
- `ylstrip/command/[devicename]/[commandname]`

### state

Receiving state updates:

- `ylcolor/state/[devicename]/[property]`
- `ylmono/state/[evicename]/[property]`
- `ylstrip/state/[devicename]/[property]`


## common properties for all models

- `[type]/state/[devicename]/name` device name
- `[type]/state/[devicename]/model` one of `mono`/`color`/`strip`
- `[type]/state/[devicename]/location` device address (example: `yeelight://192.168.2.7:55443`)
- `[type]/state/[devicename]/power` `on` or `off`
- `[type]/state/[devicename]/brightness` 1-100

## ylcolor

- color_temp 1700-6500
- rgb R,G,B (0-255)
- start_animation [name]
  - built in animations: strobe
- start_flow
- stop_flow

### commands
