package exporter

func Export() {
	flags := getParsedFlags()

	credentials := resolveCredentials(flags)
	
	if err := createDirIfNotExists(flags); err != nil {
		panic(err)
	}
}
