# ABN
Look up abn details using abn number.

```bash
go get github.com/Harry-3d/ABN 
```
```go
	// You can obtain a GUID from https://abr.business.gov.au/Tools/WebServicesAgreement
	// Set the GUID
	abn.SetGUID("my-guid-key-here")

	// Use the ABN number to look up details.
	json := abn.LookUpABN("31079595467")

	//Result
	fmt.Println(json)
	fmt.Println(json["EntityName"])
```
