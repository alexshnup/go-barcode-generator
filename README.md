# go-barcode-generator

Barcode Generator for output on the e-ink Sony Reader PRS-505. For Use With
(https://github.com/alexshnup/sony_prs-505_remote_show)


\*\*Attention!!!\*\* \*This code is write output to a disk device with low level I/O, specifically designed for the reception of the image in a specific format.  If you make a mistake with the name of the device, you may lose your data. Please check the drive name.\*

# Run

Let's find the device name disk with size of 1MB.
```
fdisk -l
```
We find a similar line in the output:
```
Disk /dev/sde: 1 MB, 1048576 bytes
```
Now we know that we must send the image to "/dev/sde"

# Ean Barcode
```
go run main.go {DISK DRIVE HERE} ean 123456789012
```

# QR Barcode
```
go run main.go {DISK DRIVE HERE} qr "1234567890 test"
```
