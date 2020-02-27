# cvr100u
读取身份证证件信息，华视电子-cvr100u
CVR100U is an physical device, you can use it read the Chinese's ID card infomation, more detail of this device can get from this [website](http://www.chinaidcard.com/)

## Installation
go get github.com/zogyi/cvr100u

## Usage

```go
conn := device.Connector{IsX64: true}
```

##Notice
The DLLs come from the ChinaVision's website, please download the latest files and replace them