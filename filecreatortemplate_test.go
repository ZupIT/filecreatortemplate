/*
 * Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filemanagertemplate

import (
    "testing"
    "os"
	"github.com/gobuffalo/packr/v2"
)

const  template = "templateExample.kt"
const templateLocation = "../filecreatortemplate/"

func TestFileCreatorTemplateSucces(t *testing.T) {
    
    var newFilePath = "./testaux/temp"
    var newFileName = "TemplateCreatedExample"
    var m = make(map[string]string)
    m["packageName"] = "templateExample"
    m["className"] = "Example"
    
    CreatFile(newFilePath, newFileName, "kt" ,templateLocation,template,m,os.Stdout)
 
    if getFileCreatedContet() != getTemplateResultExpectedContet(){
        t.Errorf("The template created is not equals to the expected")
    }
    os.Remove("../filecreatortemplate/testaux/temp/TemplateCreatedExample.kt")
}

func TestFileCreatorTemplateErrorFileAlreadyExist(t *testing.T) {
    
    var newFilePath = "./testaux/temp"
    var newFileName = "TemplateCreatedExample"
    var m = make(map[string]string)
    m["packageName"] = "templateExample"
    m["className"] = "Example"
    
    CreatFile(newFilePath, newFileName, "kt" ,templateLocation, template,m,os.Stdout)
    defer func(){
        recover()
        os.Remove("../filecreatortemplate/testaux/temp/TemplateCreatedExample.kt")

    }()
    CreatFile(newFilePath, newFileName, "kt" ,templateLocation, template,m,os.Stdout)
    t.Errorf("should panic")
}

func TestFileCreatorTemplateErrorTemplateNotFound(t *testing.T) {
    
    var newFilePath = "./testaux/temp"
    var newFileName = "TemplateCreatedExample"
    var m = make(map[string]string)
    m["packageName"] = "templateExample"
    m["className"] = "Example"

    defer func(){
        recover()
        os.Remove("../filecreatortemplate/testaux/temp/TemplateCreatedExample.kt")

    }()
    CreatFile(newFilePath, newFileName, "kt" ,templateLocation, template+"/somethingToBroke",m,os.Stdout)
    t.Errorf("should panic")
}

func getFileCreatedContet() string{
    box :=  packr.New("fileBox", "./temp")
    data, _ := box.FindString("./testaux/temp/TemplateCreatedExample.kt")
    var templateContent = string(data)
    return templateContent
}

func getTemplateResultExpectedContet() string{
    box :=  packr.New("fileBox", "./temp")
    data, _ := box.FindString("./templateExampleExpected.kt")
    var templateContent = string(data)

    return templateContent
}