package types

type TableMeta struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	DataFile   string `json:"data_file"`
	ConfigFile string `json:"config_file"`
}