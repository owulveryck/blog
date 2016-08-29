---
author: Olivier Wulveryck
date: 2016-08-29T21:58:17+02:00
description: A first geek interaction between my  raspberry pi 3 and my weather station
draft: true
tags:
- raspberry pi
title: Getting weather data from the station to the raspberry
topics:
- topic 1
type: post
---


I have plugged my weather information "oregon scientific" model "RMS 300" into my raspberry pi 3

what `dmesg` tells me is simply

```shell
[ 2256.877522] usb 1-1.4: new low-speed USB device number 5 using dwc_otg
[ 2256.984860] usb 1-1.4: New USB device found, idVendor=0fde, idProduct=ca01
[ 2256.984881] usb 1-1.4: New USB device strings: Mfr=0, Product=1, SerialNumber=0
[ 2256.984894] usb 1-1.4: Product:  
[ 2256.992719] hid-generic 0003:0FDE:CA01.0002: hiddev0,hidraw0: USB HID v1.10 Device [ ] on usb-3f980000.usb-1.4/input0
```

I have a pseudo device listed here: `crw------- 1 root root 180, 96 Aug 29 19:55 /dev/usb/hiddev0` so let's dig into it.

## Giving access: `udev`

The first thing to do is to allow access to my usb device so I won't need to run any program as root.
By default the `pi` user belongs to a bunch of groups. One of those is called `plugdev`.
It is the one I will use for my experiment.

### Get information about my Device

 `~ find /dev/bus/usb/ '!' -type d -mmin -5`
<pre>
 /dev/bus/usb/001/012
</pre>
 `~ udevadm info /dev/bus/usb/001/012`

<pre>
 P: /devices/platform/soc/3f980000.usb/usb1/1-1/1-1.3
 N: bus/usb/001/012
 E: BUSNUM=001
 E: DEVNAME=/dev/bus/usb/001/012
 E: DEVNUM=012
 E: DEVPATH=/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.3
 E: DEVTYPE=usb_device
 E: DRIVER=usb
 E: ID_BUS=usb
 E: ID_MODEL_ENC=\x20
 E: ID_MODEL_FROM_DATABASE=WMRS200 weather station
 E: ID_MODEL_ID=ca01
 E: ID_REVISION=0302
 E: ID_SERIAL=0fde_
 E: ID_USB_INTERFACES=:030000:
 E: ID_VENDOR=0fde
 E: ID_VENDOR_ENC=0fde
 E: ID_VENDOR_FROM_DATABASE=Oregon Scientific
 E: ID_VENDOR_ID=0fde
 E: MAJOR=189
 E: MINOR=11
 E: PRODUCT=fde/ca01/302
 E: SUBSYSTEM=__usb__
 E: TYPE=0/0/0
 E: USEC_INITIALIZED=5929384
 </pre>

  That is the correct device. I will note the vendor ID and the product ID


Let's create the udev file using the previous informations about the idVendor and the idProduct

```shell
cat << EOF > /etc/udev/rules.d/50-weather-station.rules
# Weather Station
SUBSYSTEM=="usb", ATTRS{idVendor}=="0fde", ATTRS{idProduct}=="ca01", MODE="0660", GROUP="plugdev"
EOF
```

And finally restart udev with `sudo /etc/init.d/udev restart`
