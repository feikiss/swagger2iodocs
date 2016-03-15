package model

type Extensions map[string]interface{}

type apiconfig struct {

	name string `json:"name"`
}

type apiconfigs struct  {
	//ApiVendorName:apiconfig
	configMap map[string]*apiconfig
}

type Iodoc struct {

	Name string `json:",omitempty"`
	Protocol string `json:"protocol,omitempty"`
	BasePath string `json:"basePath,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	PublicPath string `json:"publicPath,omitempty"`
	PrivatePath string `json:"privatePath,omitempty"`
	Version string `json:"version,omitempty"`
	Resources map[string]*Resource `json:"resources,omitempty"`
	Schemas map[string]*schema `json:"schemas,omitempty"`
}

func New() Iodoc  {
	iodoc := Iodoc{
		BasePath:"1324",
		Name:"hello",
		Protocol:"Rest",
		PublicPath:"/v2",
		Version:"2.0",
		Resources:make(map[string]*Resource),
		}

	//Resource{
	//	"":ResourceProps{
	//		Methods:make(map[string]MethodProps),
	//	},

	return iodoc
}

type schema struct {

	Properties map[string]*schemaProp `json:"properties,omitempty"` //Key: schemaProp name.

}

type schemaProp struct {
	DefaultProp string `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
	Required string `json:"required,omitempty"`
	Title string `json:"title,omitempty"`
	TypeProp string `json:"type,omitempty"`

}

type Resource struct {
	Methods map[string]*MethodProps `json:"methods,omitempty"`
}

//type Method map[string]MethodProps

type MethodProps struct  {
	Description string `json:"description,omitempty"`
	HttpMethod string `json:"httpMethod,omitempty"`
	Name string `json:"name,omitempty"`
	Parameters map[string]*ParameterProps `json:"parameters,omitempty"`
	Path string `json:"path"`
}


type ParameterProps struct {
	Default string `json:default,omitempty`
	Description     string  `json:"description,omitempty"`
	Title string `json:"title,omitempty"`
	Required bool    `json:"required,omitempty"` // when in == "query" || "formData"
	ParamType string `json:"type,omitempty"`
	Location string `json:"location,omitempty"`
}