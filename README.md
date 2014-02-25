# Process Vendor Payment Report
Process the "Export Sold Items" CSV from ShopKeep to display a report containing how much each supplier sold and how much they are owed.

## Source Install
1. Setup [Go](http://golang.org/doc/install)
2. Clone this repository
    - `git clone git://github.com/JFMarket/process-vendor-payment-report.git`
3. Run the server
    - `cd process-vendor-payment-report`
    - `go get`
    - `go run jfm-process-vendor-payment-report.go`
4. Follow the instructions at <http://localhost:3000/>

## Example Report
Name        | Total   | Fee(%) | Final
----------- | ------- | ------ | ------
Farmer John | $100.00 | 20     | $80.00
Farmer Su   | $60.00  | 20     | $48.00
