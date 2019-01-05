package metaDataConfig

//structure for yml file
type MetaDataConfig struct{
    Title string
    Version string
    Maintainers []map[string]string
    Company string
    Website string
    Source string
    License string
    Description string
}