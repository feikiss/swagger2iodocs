package main

import (
	"fmt"
	"io/ioutil"
	"github.com/go-swagger/go-swagger/spec"
	"encoding/json"
	"swaggerconverter/model"
	"strings"
)

func main()  {

	b,err := ioutil.ReadFile("data/TIBCO Cloud Integration PetStore Sample.json")
	if err != nil {
		fmt.Println(err)
	}

	doc,docerr := spec.New(json.RawMessage(b), "")
	fmt.Println("All definations:",doc.AllDefinitions())
	definations := doc.AllDefinitions()
	fmt.Println(definations[0].Name)
	fmt.Println(definations[0].Ref)
	fmt.Println(definations[0].Schema)
	if docerr != nil {
		fmt.Println(docerr)
		panic(docerr)
	}

	iodoc := convertSwaggerToIodoc(doc)
	content,_ := json.Marshal(iodoc)
	fmt.Println(string(content))

	fileName := strings.Replace(iodoc.Name," ","",-1)

	configContent := fmt.Sprintf(`{
	"%s": {
		"name": "%s"
	}
}`,
		fileName,iodoc.Name)

	if fileErr := ioutil.WriteFile("iodoc/apiconfig.json",[]byte(configContent),0644);fileErr!=nil{
		fmt.Println(fileErr)
		panic(fileErr)
	}

	fileerr := ioutil.WriteFile("iodoc/"+fileName+".json",content,0644)
	if fileerr != nil {
		fmt.Println(fileerr)
	}
	fmt.Printf("The files %s and %s are generated successfully.\n","apiconfig.json",fileName+".json")
}

func convertSwaggerToIodoc(doc *spec.Document) *model.Iodoc {
	iodoc := new(model.Iodoc)
	spec := doc.Spec()
	publicPrefix := "https://"
	if len(spec.Schemes) > 0{
		publicPrefix = spec.Schemes[0] + "://"
	}
	iodoc.BasePath = publicPrefix + spec.Host
	iodoc.PublicPath = spec.BasePath
	iodoc.PrivatePath = spec.BasePath
	iodoc.Protocol = "rest"
	iodoc.Name = spec.Info.Title

	iodoc.Headers = func() map[string]string {
		headMap := make(map[string]string)
		if len(spec.Consumes) > 0 {
			headMap["Content-Type"] = spec.Consumes[0]
		}
		return headMap
	}()

	iodoc.Version = spec.Info.Version
	iodoc.Resources = map[string]*model.Resource{
		"UnitMethod":func() *model.Resource{
			resource := &model.Resource{}
			resource.Methods = getMethodPropsMap(doc)
			return resource
		}(),
	}

	return iodoc
}

func getMethodPropsMap(doc *spec.Document) map[string]*model.MethodProps {
	methodPropsMap := make(map[string]*model.MethodProps)

	for method,operations := range doc.Operations() {
		for path,operation := range operations {

			//fmt.Println(opeKey,"-:-",opeValue)
			var methodName string  //The alias of this method.
			if len(operation.ID) > 0 {
				methodName = operation.ID
			}else {
				methodName = method + strings.Replace(path,"/","_",-1)
			}
			methodProps := model.MethodProps{
				Name : methodName,
				Description : operation.Description,
				HttpMethod : method,
				Path : path,
			}

			//fmt.Println("operation:",*operation)

			parametersMap := make(map[string]*model.ParameterProps)

			//get the parameter
			for _, swaggerParameter := range operation.Parameters{
				paramProps := &model.ParameterProps{
					Title:swaggerParameter.Name,
					Description:swaggerParameter.Description,
					ParamType : swaggerParameter.Type,
					Required : swaggerParameter.Required,
					Location:swaggerParameter.In,
				}
				if swaggerParameter.Default != nil {
					paramProps.Default = swaggerParameter.Default.(string)
				}
				parametersMap[paramProps.Title] = paramProps

				methodProps.Parameters = parametersMap
				methodPropsMap[methodName] = &methodProps
			}
		}

	}
	return methodPropsMap
}

func readIodoc()  {
	iodcont,err := ioutil.ReadFile("rdio2.json")
	if err != nil {
		println(err)
	}

	iodoc2 := new(model.Iodoc)
	json.Unmarshal(iodcont,&iodoc2)
	if b,err := json.Marshal(iodoc2);err ==nil{
		fmt.Println(string(b))

		iodoc3 := new(model.Iodoc)
		json.Unmarshal(b,iodoc3)
		fmt.Println(iodoc3)
	}
}