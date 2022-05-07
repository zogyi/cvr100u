# cvr100u
读取身份证证件信息，华视电子-cvr100u

CVR100U is a hardware device which can read the info from the Chinese ID card, it's a product of the this compnay, you check it out on this website [website](http://www.chinaidcard.com/).

You can use this library to get realname, gender, nation, birthday, home address and the photo as same as the ID card showing.




## Installation
go get github.com/zogyi/cvr100u

## Usage

```go
conn := device.Connector{IsX64: true}
conn.Initial()            // initial dll into the program
conn.Authentication()     // do authentication
conn.ReadContent()        //read the content from the ID card
conn.ReadFields(ReadName) //get the realname

```

### Notice
The DLLs come from the ChinaVision's website, please download the latest files and replace them
