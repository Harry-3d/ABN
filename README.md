# ABN
Look up abn details using abn number.

	// obtain a GUID from https://abr.business.gov.au/Tools/WebServicesAgreement
	abn.SetGUID("29d06b3e-bb76-46b2-9e92-aeaba9536546")

	// Use the ABN number to look up details.
	json := abn.LookUpABN("31079595467")

	//Result
	fmt.Println(json)
	fmt.Println(json["EntityName"])
